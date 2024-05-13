package bindings

import (
	"context"

	"github.com/Southclaws/dt"
	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/ftag"
	"github.com/Southclaws/opt"
	"github.com/rs/xid"
	"github.com/samber/lo"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/asset"
	"github.com/Southclaws/storyden/app/resources/datagraph"
	"github.com/Southclaws/storyden/app/resources/datagraph/cluster_traversal"
	"github.com/Southclaws/storyden/app/resources/post"
	"github.com/Southclaws/storyden/app/services/authentication/session"
	cluster_svc "github.com/Southclaws/storyden/app/services/cluster"
	"github.com/Southclaws/storyden/app/services/cluster/cluster_visibility"
	"github.com/Southclaws/storyden/app/services/clustertree"
	"github.com/Southclaws/storyden/app/transports/openapi"
)

type Nodes struct {
	ar    account.Repository
	cs    cluster_svc.Manager
	cv    *cluster_visibility.Controller
	ctree clustertree.Graph
	ctr   cluster_traversal.Repository
}

func NewNodes(
	ar account.Repository,
	cs cluster_svc.Manager,
	cv *cluster_visibility.Controller,
	ctree clustertree.Graph,
	ctr cluster_traversal.Repository,
) Nodes {
	return Nodes{
		ar:    ar,
		cs:    cs,
		cv:    cv,
		ctree: ctree,
		ctr:   ctr,
	}
}

func (c *Nodes) NodeCreate(ctx context.Context, request openapi.NodeCreateRequestObject) (openapi.NodeCreateResponseObject, error) {
	session, err := session.GetAccountID(ctx)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	vis, err := opt.MapErr(opt.NewPtr(request.Body.Visibility), deserialiseVisibility)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	node, err := c.cs.Create(ctx,
		session,
		request.Body.Name,
		request.Body.Slug,
		request.Body.Description,
		cluster_svc.Partial{
			Content:    opt.NewPtr(request.Body.Content),
			Properties: opt.NewPtr(request.Body.Properties),
			URL:        opt.NewPtr(request.Body.Url),
			AssetsAdd:  opt.NewPtrMap(request.Body.AssetIds, deserialiseAssetIDs),
			Parent:     opt.NewPtrMap(request.Body.Parent, deserialiseClusterSlug),
			Visibility: vis,
		},
	)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeCreate200JSONResponse{
		NodeCreateOKJSONResponse: openapi.NodeCreateOKJSONResponse(serialiseCluster(node)),
	}, nil
}

func (c *Nodes) NodeList(ctx context.Context, request openapi.NodeListRequestObject) (openapi.NodeListResponseObject, error) {
	acc, err := opt.MapErr(session.GetOptAccountID(ctx), func(aid account.AccountID) (*account.Account, error) {
		return c.ar.GetByID(ctx, aid)
	})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	var cs []*datagraph.Cluster

	opts := []cluster_traversal.Filter{}

	if v := request.Params.Author; v != nil {
		opts = append(opts, cluster_traversal.WithOwner(*v))
	}

	visibilities, err := opt.MapErr(opt.NewPtr(request.Params.Visibility), deserialiseVisibilityList)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	if a, ok := acc.Get(); ok {
		// NOTE: We do not want to allow anyone to request ANY cluster that is
		// not published, but we also want to allow admins to request clusters
		// that are in review. So we need to check if the requesting account is
		// an admin and if they are not, automatically add a WithOwner filter.

		if v, ok := visibilities.Get(); ok {
			opts = append(opts, cluster_traversal.WithVisibility(v...))

			if lo.Contains(v, post.VisibilityDraft) {
				// If the result is to contain drafts, only show the account's.
				opts = append(opts, cluster_traversal.WithOwner(a.Handle))
			} else if lo.Contains(v, post.VisibilityReview) {
				// If the result is to contain clusters that are in-review, then
				// we need to check if the requesting account is an admin first.
				if !a.Admin {
					opts = append(opts, cluster_traversal.WithOwner(a.Handle))
				}
			}
		}
	} else {
		// When the request is not made by an authenticated account, we do not
		// permit any visibility other than "published".

		opts = append(opts, cluster_traversal.WithVisibility(post.VisibilityPublished))
	}

	if id := request.Params.NodeId; id != nil {
		cid, err := xid.FromString(*id)
		if err != nil {
			return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.InvalidArgument))
		}

		cs, err = c.ctr.Subtree(ctx, datagraph.ClusterID(cid), opts...)
		if err != nil {
			return nil, fault.Wrap(err, fctx.With(ctx))
		}
	} else {
		cs, err = c.ctr.Root(ctx, opts...)
		if err != nil {
			return nil, fault.Wrap(err, fctx.With(ctx))
		}
	}

	return openapi.NodeList200JSONResponse{
		NodeListOKJSONResponse: openapi.NodeListOKJSONResponse{
			Nodes: dt.Map(cs, serialiseCluster),
		},
	}, nil
}

func (c *Nodes) NodeGet(ctx context.Context, request openapi.NodeGetRequestObject) (openapi.NodeGetResponseObject, error) {
	node, err := c.cs.Get(ctx, datagraph.ClusterSlug(request.NodeSlug))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeGet200JSONResponse{
		NodeGetOKJSONResponse: openapi.NodeGetOKJSONResponse(serialiseClusterWithItems(node)),
	}, nil
}

func (c *Nodes) NodeUpdate(ctx context.Context, request openapi.NodeUpdateRequestObject) (openapi.NodeUpdateResponseObject, error) {
	node, err := c.cs.Update(ctx, datagraph.ClusterSlug(request.NodeSlug), cluster_svc.Partial{
		Name:        opt.NewPtr(request.Body.Name),
		Slug:        opt.NewPtr(request.Body.Slug),
		AssetsAdd:   opt.NewPtrMap(request.Body.AssetIds, deserialiseAssetIDs),
		URL:         opt.NewPtr(request.Body.Url),
		Description: opt.NewPtr(request.Body.Description),
		Content:     opt.NewPtr(request.Body.Content),
		Parent:      opt.NewPtrMap(request.Body.Parent, deserialiseClusterSlug),
		Properties:  opt.NewPtr(request.Body.Properties),
	})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeUpdate200JSONResponse{
		NodeUpdateOKJSONResponse: openapi.NodeUpdateOKJSONResponse(serialiseCluster(node)),
	}, nil
}

func (c *Nodes) NodeUpdateVisibility(ctx context.Context, request openapi.NodeUpdateVisibilityRequestObject) (openapi.NodeUpdateVisibilityResponseObject, error) {
	v, err := post.NewVisibility(string(request.Body.Visibility))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.InvalidArgument))
	}

	node, err := c.cv.ChangeVisibility(ctx, datagraph.ClusterSlug(request.NodeSlug), v)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeUpdateVisibility200JSONResponse{
		NodeUpdateOKJSONResponse: openapi.NodeUpdateOKJSONResponse(serialiseCluster(node)),
	}, nil
}

func (c *Nodes) NodeDelete(ctx context.Context, request openapi.NodeDeleteRequestObject) (openapi.NodeDeleteResponseObject, error) {
	destinationCluster, err := c.cs.Delete(ctx, datagraph.ClusterSlug(request.NodeSlug), cluster_svc.DeleteOptions{
		MoveTo:   opt.NewPtr((*datagraph.ClusterSlug)(request.Params.TargetNode)),
		Clusters: opt.NewPtr(request.Params.MoveChildNodes).OrZero(),
	})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeDelete200JSONResponse{
		NodeDeleteOKJSONResponse: openapi.NodeDeleteOKJSONResponse{
			Destination: opt.Map(opt.NewPtr(destinationCluster), func(in datagraph.Cluster) openapi.Node {
				return serialiseCluster(&in)
			}).Ptr(),
		},
	}, nil
}

func (c *Nodes) NodeAddAsset(ctx context.Context, request openapi.NodeAddAssetRequestObject) (openapi.NodeAddAssetResponseObject, error) {
	id := openapi.ParseID(request.AssetId)

	node, err := c.cs.Update(ctx, datagraph.ClusterSlug(request.NodeSlug), cluster_svc.Partial{
		AssetsAdd: opt.New([]asset.AssetID{id}),
	})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeAddAsset200JSONResponse{
		NodeUpdateOKJSONResponse: openapi.NodeUpdateOKJSONResponse(serialiseCluster(node)),
	}, nil
}

func (c *Nodes) NodeRemoveAsset(ctx context.Context, request openapi.NodeRemoveAssetRequestObject) (openapi.NodeRemoveAssetResponseObject, error) {
	id := openapi.ParseID(request.AssetId)

	node, err := c.cs.Update(ctx, datagraph.ClusterSlug(request.NodeSlug), cluster_svc.Partial{
		AssetsRemove: opt.New([]asset.AssetID{id}),
	})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeRemoveAsset200JSONResponse{
		NodeUpdateOKJSONResponse: openapi.NodeUpdateOKJSONResponse(serialiseCluster(node)),
	}, nil
}

func (c *Nodes) NodeAddNode(ctx context.Context, request openapi.NodeAddNodeRequestObject) (openapi.NodeAddNodeResponseObject, error) {
	node, err := c.ctree.Move(ctx, datagraph.ClusterSlug(request.NodeSlugChild), datagraph.ClusterSlug(request.NodeSlug))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeAddNode200JSONResponse{
		NodeAddChildOKJSONResponse: openapi.NodeAddChildOKJSONResponse(serialiseCluster(node)),
	}, nil
}

func (c *Nodes) NodeRemoveNode(ctx context.Context, request openapi.NodeRemoveNodeRequestObject) (openapi.NodeRemoveNodeResponseObject, error) {
	node, err := c.ctree.Sever(ctx, datagraph.ClusterSlug(request.NodeSlugChild), datagraph.ClusterSlug(request.NodeSlug))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeRemoveNode200JSONResponse{
		NodeRemoveChildOKJSONResponse: openapi.NodeRemoveChildOKJSONResponse(serialiseCluster(node)),
	}, nil
}

func serialiseCluster(in *datagraph.Cluster) openapi.Node {
	return openapi.Node{
		Id:          in.ID.String(),
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
		Name:        in.Name,
		Slug:        in.Slug,
		Assets:      dt.Map(in.Assets, serialiseAssetReference),
		Link:        opt.Map(in.Links.Latest(), serialiseLink).Ptr(),
		Description: in.Description,
		Content:     in.Content.Ptr(),
		Owner:       serialiseProfileReference(in.Owner),
		Parent:      opt.Map(in.Parent, serialiseCluster).Ptr(),
		Visibility:  serialiseVisibility(in.Visibility),
		Properties:  in.Properties,
	}
}

func serialiseClusterWithItems(in *datagraph.Cluster) openapi.NodeWithChildren {
	return openapi.NodeWithChildren{
		Id:          in.ID.String(),
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
		Name:        in.Name,
		Slug:        in.Slug,
		Assets:      dt.Map(in.Assets, serialiseAssetReference),
		Link:        opt.Map(in.Links.Latest(), serialiseLink).Ptr(),
		Description: in.Description,
		Content:     in.Content.Ptr(),
		Owner:       serialiseProfileReference(in.Owner),
		Parent:      opt.Map(in.Parent, serialiseCluster).Ptr(),
		Visibility:  serialiseVisibility(in.Visibility),
		Properties:  in.Properties,
		Children:    dt.Map(in.Clusters, serialiseCluster),
	}
}

func deserialiseClusterSlug(in string) datagraph.ClusterSlug {
	return datagraph.ClusterSlug(in)
}
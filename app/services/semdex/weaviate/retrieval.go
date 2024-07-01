package weaviate

import (
	"context"
	"encoding/json"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/storyden/app/resources/datagraph"
	"github.com/weaviate/weaviate/entities/models"
)

func (o *weaviateSemdexer) GetAll(ctx context.Context) (datagraph.NodeReferenceList, error) {
	objects, err := o.wc.Data().
		ObjectsGetter().
		WithClassName(o.cn.String()).
		Do(ctx)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	invalid := []*models.Object{}
	refs := []*datagraph.NodeReference{}

	for _, m := range objects {
		o, err := mapWeaviateObject(m)
		if err != nil {
			invalid = append(invalid, m)
			continue
		}

		refs = append(refs, o)
	}

	for _, m := range invalid {
		if err := o.wc.Data().Deleter().WithID(m.ID.String()).Do(ctx); err != nil {
			return nil, fault.Wrap(err, fctx.With(ctx))
		}
	}

	return datagraph.NodeReferenceList(refs), nil
}

func mapWeaviateObject(o *models.Object) (*datagraph.NodeReference, error) {
	wo, err := unmarshalWeaviateObject(o.Properties)
	if err != nil {
		return nil, err
	}

	ref, err := mapToNodeReference(*wo)
	if err != nil {
		return nil, err
	}

	return ref, nil
}

func unmarshalWeaviateObject(p models.PropertySchema) (*WeaviateObject, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	wo := WeaviateObject{}
	if err := json.Unmarshal(b, &wo); err != nil {
		return nil, err
	}

	return &wo, nil
}

package datagraph_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"

	"github.com/Southclaws/dt"
	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/seed"
	"github.com/Southclaws/storyden/app/transports/openapi/bindings"
	"github.com/Southclaws/storyden/internal/integration"
	"github.com/Southclaws/storyden/internal/integration/e2e"
	"github.com/Southclaws/storyden/internal/openapi"
)

func TestDatagraphHappyPath(t *testing.T) {
	t.Parallel()

	integration.Test(t, nil, e2e.Setup(), fx.Invoke(func(
		lc fx.Lifecycle,
		ctx context.Context,
		cl *openapi.ClientWithResponses,
		cj *bindings.CookieJar,
		ar account.Repository,
	) {
		lc.Append(fx.StartHook(func() {
			r := require.New(t)
			a := assert.New(t)

			ctx, acc := e2e.WithAccount(ctx, ar, seed.Account_001_Odin)

			name1 := "test-cluster-1"
			slug1 := name1 + uuid.NewString()
			clus1, err := cl.ClusterCreateWithResponse(ctx, openapi.ClusterInitialProps{
				Name:        name1,
				Slug:        slug1,
				Description: "testing clusters api",
			}, e2e.WithSession(ctx, cj))
			r.NoError(err)
			r.NotNil(clus1)
			r.Equal(200, clus1.StatusCode())

			a.Equal(name1, clus1.JSON200.Name)
			a.Equal(slug1, clus1.JSON200.Slug)
			a.Equal("testing clusters api", clus1.JSON200.Description)
			a.Equal(acc.ID.String(), string(clus1.JSON200.Owner.Id))

			// Add a child cluster

			name2 := "test-cluster-2"
			slug2 := name2 + uuid.NewString()
			clus2, err := cl.ClusterCreateWithResponse(ctx, openapi.ClusterInitialProps{
				Name:        name2,
				Slug:        slug2,
				Description: "testing clusters children",
			}, e2e.WithSession(ctx, cj))
			r.NoError(err)
			r.NotNil(clus2)
			r.Equal(200, clus2.StatusCode())

			cadd, err := cl.ClusterAddClusterWithResponse(ctx, slug1, slug2, e2e.WithSession(ctx, cj))
			r.NoError(err)
			r.NotNil(cadd)
			r.Equal(200, cadd.StatusCode())
			r.Equal(clus1.JSON200.Id, cadd.JSON200.Id)

			// Add another child to this child
			// clus1
			// |- clus2
			//    |- clus3

			name3 := "test-cluster-3"
			slug3 := name3 + uuid.NewString()
			clus3, err := cl.ClusterCreateWithResponse(ctx, openapi.ClusterInitialProps{
				Name:        name3,
				Slug:        slug3,
				Description: "testing clusters children",
			}, e2e.WithSession(ctx, cj))
			r.NoError(err)
			r.NotNil(clus3)
			r.Equal(200, clus3.StatusCode())

			cadd, err = cl.ClusterAddClusterWithResponse(ctx, slug2, slug3, e2e.WithSession(ctx, cj))
			r.NoError(err)
			r.NotNil(cadd)
			r.Equal(200, cadd.StatusCode())
			r.Equal(clus2.JSON200.Id, cadd.JSON200.Id)

			// Create item

			itemname1 := "test-item-1-" + uuid.NewString()
			itemslug1 := itemname1
			item1, err := cl.ItemCreateWithResponse(ctx, openapi.ItemInitialProps{
				Name:        itemname1,
				Slug:        itemslug1,
				Description: "testing items api",
			}, e2e.WithSession(ctx, cj))
			r.NoError(err)
			r.NotNil(item1)
			r.Equal(200, item1.StatusCode())

			// Add item to clus1

			clus1added, err := cl.ClusterAddItemWithResponse(ctx, clus1.JSON200.Slug, item1.JSON200.Slug, e2e.WithSession(ctx, cj))
			r.NoError(err)
			r.NotNil(clus1added)
			r.Equal(200, clus1added.StatusCode())

			// Get clus1

			clus1get, err := cl.ClusterGetWithResponse(ctx, clus1.JSON200.Slug)
			r.NoError(err)
			r.NotNil(clus1get)
			r.Equal(200, clus1get.StatusCode())

			itemids := dt.Map(clus1get.JSON200.Items, func(i openapi.Item) string { return i.Id })
			a.Contains(itemids, item1.JSON200.Id)
		}))
	}))
}

package node_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/seed"
	"github.com/Southclaws/storyden/app/transports/openapi"
	"github.com/Southclaws/storyden/app/transports/openapi/bindings"
	"github.com/Southclaws/storyden/internal/integration"
	"github.com/Southclaws/storyden/internal/integration/e2e"
	"github.com/Southclaws/storyden/tests"
)

func TestNodesHappyPath(t *testing.T) {
	t.Parallel()

	integration.Test(t, nil, e2e.Setup(), fx.Invoke(func(
		lc fx.Lifecycle,
		ctx context.Context,
		cl *openapi.ClientWithResponses,
		cj *bindings.CookieJar,
		ar account.Repository,
	) {
		lc.Append(fx.StartHook(func() {
			a := assert.New(t)

			ctx, acc := e2e.WithAccount(ctx, ar, seed.Account_001_Odin)

			visibility := openapi.Published

			name1 := "test-node-1"
			slug1 := name1 + uuid.NewString()
			content1 := "# Nodes\n\nRich text content."
			// iurl1 := "https://picsum.photos/200/200"
			url1 := "https://southcla.ws"
			node1, err := cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{
				Name:        name1,
				Slug:        slug1,
				Description: "testing nodes api",
				Content:     &content1,
				Url:         &url1,
				Visibility:  &visibility, // Admin account can post directly to published
			}, e2e.WithSession(ctx, cj))
			tests.Ok(t, err, node1)

			a.Equal(name1, node1.JSON200.Name)
			a.Equal(slug1, node1.JSON200.Slug)
			a.Equal("testing nodes api", node1.JSON200.Description)
			a.Equal(content1, *node1.JSON200.Content)
			a.Equal(url1, node1.JSON200.Link.Url)
			a.Equal(acc.ID.String(), string(node1.JSON200.Owner.Id))

			// Get the one just created

			node1get, err := cl.NodeGetWithResponse(ctx, slug1)
			tests.Ok(t, err, node1get)

			a.Equal(name1, node1get.JSON200.Name)
			a.Equal(slug1, node1get.JSON200.Slug)
			a.Equal("testing nodes api", node1get.JSON200.Description)
			a.Equal(acc.ID.String(), string(node1get.JSON200.Owner.Id))

			// Update the one just created

			name1 = "test-node-1-UPDATED"
			slug1 = name1 + uuid.NewString()
			desc1 := "a new description"
			cont1 := "# New content"
			// iurl1 = "https://picsum.photos/500/500"
			url1 = "https://cla.ws"
			prop1 := any(map[string]any{
				"key": "value",
			})
			node1update, err := cl.NodeUpdateWithResponse(ctx, node1.JSON200.Slug, openapi.NodeMutableProps{
				Name:        &name1,
				Slug:        &slug1,
				Description: &desc1,
				Content:     &cont1,
				Url:         &url1,
				Properties:  &prop1,
			}, e2e.WithSession(ctx, cj))
			tests.Ok(t, err, node1update)

			a.Equal(name1, node1update.JSON200.Name)
			a.Equal(slug1, node1update.JSON200.Slug)
			a.Equal(desc1, node1update.JSON200.Description)
			a.Equal(cont1, *node1update.JSON200.Content)
			a.Equal(url1, node1update.JSON200.Link.Url)
			a.Equal(prop1, node1update.JSON200.Properties)

		}))
	}))
}

func TestNodesErrors(t *testing.T) {
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

			ctx, _ := e2e.WithAccount(ctx, ar, seed.Account_001_Odin)

			get404, err := cl.NodeGetWithResponse(ctx, "nonexistent")
			r.NoError(err)
			r.NotNil(get404)
			a.Equal(http.StatusNotFound, get404.StatusCode())

			update403, err := cl.NodeUpdateWithResponse(ctx, "nonexistent", openapi.NodeMutableProps{})
			r.NoError(err)
			r.NotNil(update403)
			a.Equal(http.StatusForbidden, update403.StatusCode())

			update404, err := cl.NodeUpdateWithResponse(ctx, "nonexistent", openapi.NodeMutableProps{}, e2e.WithSession(ctx, cj))
			r.NoError(err)
			r.NotNil(update404)
			a.Equal(http.StatusNotFound, update404.StatusCode())
		}))
	}))
}
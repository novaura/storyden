package library_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Southclaws/dt"
	"github.com/Southclaws/opt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"

	"github.com/Southclaws/storyden/app/resources/account/account_writer"
	"github.com/Southclaws/storyden/app/resources/seed"
	"github.com/Southclaws/storyden/app/transports/http/middleware/session"
	"github.com/Southclaws/storyden/app/transports/http/openapi"
	"github.com/Southclaws/storyden/internal/integration"
	"github.com/Southclaws/storyden/internal/integration/e2e"
	"github.com/Southclaws/storyden/tests"
)

func TestNodesVisibilityRules_Draft(t *testing.T) {
	t.Parallel()

	integration.Test(t, nil, e2e.Setup(), fx.Invoke(func(
		lc fx.Lifecycle,
		ctx context.Context,
		cl *openapi.ClientWithResponses,
		cj *session.Jar,
		aw account_writer.Writer,
	) {
		lc.Append(fx.StartHook(func() {
			ctxAdmin, _ := e2e.WithAccount(ctx, aw, seed.Account_001_Odin)
			ctxAuthor, _ := e2e.WithAccount(ctx, aw, seed.Account_003_Baldur)
			ctxRando, _ := e2e.WithAccount(ctx, aw, seed.Account_004_Loki)

			draft := openapi.Draft
			unlisted := openapi.Unlisted
			review := openapi.Review
			published := openapi.Published

			parentNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &draft}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)

			t.Run("draft_child_succeeds", func(t *testing.T) {
				t.Parallel()

				draftNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &draft}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, draftNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, draftNode.JSON200.Id)
			})

			t.Run("unlisted_child_fails", func(t *testing.T) {
				t.Parallel()

				unlistedNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &unlisted}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, unlistedNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusBadRequest)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, unlistedNode.JSON200.Id)
			})

			t.Run("review_child_fails", func(t *testing.T) {
				t.Parallel()

				reviewNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &review}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, reviewNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusBadRequest)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, reviewNode.JSON200.Id)
			})

			t.Run("published_child_fails", func(t *testing.T) {
				t.Parallel()

				publishedNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &published}, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, publishedNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusBadRequest)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.Contains(t, ids, publishedNode.JSON200.Id)
			})
		}))
	}))
}

func TestNodesVisibilityRules_Unlisted(t *testing.T) {
	t.Parallel()

	integration.Test(t, nil, e2e.Setup(), fx.Invoke(func(
		lc fx.Lifecycle,
		ctx context.Context,
		cl *openapi.ClientWithResponses,
		cj *session.Jar,
		aw account_writer.Writer,
	) {
		lc.Append(fx.StartHook(func() {
			ctxAdmin, _ := e2e.WithAccount(ctx, aw, seed.Account_001_Odin)
			ctxAuthor, _ := e2e.WithAccount(ctx, aw, seed.Account_003_Baldur)
			ctxRando, _ := e2e.WithAccount(ctx, aw, seed.Account_004_Loki)

			draft := openapi.Draft
			unlisted := openapi.Unlisted
			review := openapi.Review
			published := openapi.Published

			parentNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &unlisted}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)

			t.Run("draft_child_fails", func(t *testing.T) {
				t.Parallel()

				draftNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &draft}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, draftNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusBadRequest)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, draftNode.JSON200.Id)
			})

			t.Run("unlisted_child_succeeds", func(t *testing.T) {
				t.Parallel()

				unlistedNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &unlisted}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, unlistedNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, unlistedNode.JSON200.Id)
			})

			t.Run("review_child_fails", func(t *testing.T) {
				t.Parallel()

				reviewNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &review}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, reviewNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusBadRequest)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, reviewNode.JSON200.Id)
			})

			t.Run("published_child_fails", func(t *testing.T) {
				t.Parallel()

				publishedNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &published}, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, publishedNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusBadRequest)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.Contains(t, ids, publishedNode.JSON200.Id)
			})
		}))
	}))
}

func TestNodesVisibilityRules_Review(t *testing.T) {
	t.Parallel()

	integration.Test(t, nil, e2e.Setup(), fx.Invoke(func(
		lc fx.Lifecycle,
		ctx context.Context,
		cl *openapi.ClientWithResponses,
		cj *session.Jar,
		aw account_writer.Writer,
	) {
		lc.Append(fx.StartHook(func() {
			ctxAdmin, _ := e2e.WithAccount(ctx, aw, seed.Account_001_Odin)
			ctxAuthor, _ := e2e.WithAccount(ctx, aw, seed.Account_003_Baldur)
			ctxRando, _ := e2e.WithAccount(ctx, aw, seed.Account_004_Loki)

			draft := openapi.Draft
			unlisted := openapi.Unlisted
			review := openapi.Review
			published := openapi.Published

			parentNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &review}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)

			t.Run("draft_child_succeeds", func(t *testing.T) {
				t.Parallel()

				draftNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &draft}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, draftNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, draftNode.JSON200.Id)
			})

			t.Run("unlisted_child_fails", func(t *testing.T) {
				t.Parallel()

				unlistedNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &unlisted}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, unlistedNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusBadRequest)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, unlistedNode.JSON200.Id)
			})

			t.Run("review_child_succeeds", func(t *testing.T) {
				t.Parallel()

				reviewNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &review}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, reviewNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, reviewNode.JSON200.Id)
			})

			t.Run("published_child_fails", func(t *testing.T) {
				t.Parallel()

				publishedNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &published}, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, publishedNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusBadRequest)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.Contains(t, ids, publishedNode.JSON200.Id)
			})
		}))
	}))
}

func TestNodesVisibilityRules_Published(t *testing.T) {
	t.Parallel()

	integration.Test(t, nil, e2e.Setup(), fx.Invoke(func(
		lc fx.Lifecycle,
		ctx context.Context,
		cl *openapi.ClientWithResponses,
		cj *session.Jar,
		aw account_writer.Writer,
	) {
		lc.Append(fx.StartHook(func() {
			ctxAdmin, _ := e2e.WithAccount(ctx, aw, seed.Account_001_Odin)
			ctxAuthor, _ := e2e.WithAccount(ctx, aw, seed.Account_003_Baldur)
			ctxRando, _ := e2e.WithAccount(ctx, aw, seed.Account_004_Loki)

			draft := openapi.Draft
			unlisted := openapi.Unlisted
			review := openapi.Review
			published := openapi.Published

			parentNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n1", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &published}, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)

			t.Run("draft_child_succeeds", func(t *testing.T) {
				t.Parallel()

				draftNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &draft}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, draftNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, draftNode.JSON200.Id)
			})

			t.Run("unlisted_child_fails", func(t *testing.T) {
				t.Parallel()

				unlistedNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &unlisted}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, unlistedNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusBadRequest)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, unlistedNode.JSON200.Id)
			})

			t.Run("review_child_succeeds", func(t *testing.T) {
				t.Parallel()

				reviewNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &review}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, reviewNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, reviewNode.JSON200.Id)
			})

			t.Run("published_child_succeeds", func(t *testing.T) {
				t.Parallel()

				publishedNode := tests.AssertRequest(cl.NodeCreateWithResponse(ctx, openapi.NodeInitialProps{Name: "n2", Slug: opt.New(uuid.NewString()).Ptr(), Visibility: &draft}, e2e.WithSession(ctxAuthor, cj)))(t, http.StatusOK)
				tests.AssertRequest(cl.NodeAddNodeWithResponse(ctx, parentNode.JSON200.Slug, publishedNode.JSON200.Slug, e2e.WithSession(ctxAdmin, cj)))(t, http.StatusOK)

				list := tests.AssertRequest(cl.NodeListWithResponse(ctx, &openapi.NodeListParams{}, e2e.WithSession(ctxRando, cj)))(t, http.StatusOK)
				ids := nodeIDs(list.JSON200.Nodes)
				assert.NotContains(t, ids, publishedNode.JSON200.Id)
			})
		}))
	}))
}

func nodeIDs(nodes []openapi.NodeWithChildren) []string {
	return dt.Map(nodes, func(c openapi.NodeWithChildren) string { return c.Id })
}

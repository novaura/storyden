package thread

import (
	"context"
	"time"

	"github.com/Southclaws/dt"
	"github.com/Southclaws/opt"
	"github.com/rs/xid"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/asset"
	"github.com/Southclaws/storyden/app/resources/datagraph"
	"github.com/Southclaws/storyden/app/resources/pagination"
	"github.com/Southclaws/storyden/app/resources/post"
	"github.com/Southclaws/storyden/app/resources/post/category"
	"github.com/Southclaws/storyden/app/resources/tag/tag_ref"
	"github.com/Southclaws/storyden/app/resources/visibility"
	"github.com/Southclaws/storyden/internal/ent"
	ent_account "github.com/Southclaws/storyden/internal/ent/account"
	ent_category "github.com/Southclaws/storyden/internal/ent/category"
	ent_post "github.com/Southclaws/storyden/internal/ent/post"
	ent_tag "github.com/Southclaws/storyden/internal/ent/tag"
)

type Result struct {
	PageSize    int
	Results     int
	TotalPages  int
	CurrentPage int
	NextPage    opt.Optional[int]
	Threads     []*Thread
}

type Repository interface {
	Create(
		ctx context.Context,
		title string,
		authorID account.AccountID,
		categoryID category.CategoryID,
		opts ...Option,
	) (*Thread, error)

	Update(ctx context.Context, id post.ID, opts ...Option) (*Thread, error)

	// List is used for listing threads or filtering with basic queries. It's
	// not sufficient for full scale text-based or semantic search however.
	List(ctx context.Context,
		page int,
		size int,
		accountID opt.Optional[account.AccountID],
		opts ...Query,
	) (*Result, error)

	// GetPostCounts(ctx context.Context) (map[string]int, error)

	Get(ctx context.Context, threadID post.ID, pageParams pagination.Parameters, accountID opt.Optional[account.AccountID]) (*Thread, error)

	Delete(ctx context.Context, id post.ID) error
}

type Option func(*ent.PostMutation)

func WithID(id post.ID) Option {
	return func(m *ent.PostMutation) {
		m.SetID(xid.ID(id))
	}
}

func WithIndexed() Option {
	return func(m *ent.PostMutation) {
		m.SetIndexedAt(time.Now())
	}
}

func WithTitle(v string) Option {
	return func(pm *ent.PostMutation) {
		pm.SetTitle(v)
	}
}

func WithContent(v datagraph.Content) Option {
	return func(pm *ent.PostMutation) {
		pm.SetBody(v.HTML())
		pm.SetShort(v.Short())
	}
}

func WithCategory(v xid.ID) Option {
	return func(pm *ent.PostMutation) {
		pm.SetCategoryID(v)
	}
}

func WithVisibility(v visibility.Visibility) Option {
	return func(pm *ent.PostMutation) {
		pm.SetVisibility(ent_post.Visibility(v.String()))
	}
}

func WithMeta(meta map[string]any) Option {
	return func(m *ent.PostMutation) {
		m.SetMetadata(meta)
	}
}

func WithAssets(a []asset.AssetID) Option {
	return func(m *ent.PostMutation) {
		m.AddAssetIDs(a...)
	}
}

func WithLink(id xid.ID) Option {
	return func(pm *ent.PostMutation) {
		pm.SetLinkID(id)
	}
}

func WithContentLinks(ids ...xid.ID) Option {
	return func(pm *ent.PostMutation) {
		pm.AddContentLinkIDs(ids...)
	}
}

func WithTagsAdd(refs ...tag_ref.ID) Option {
	ids := dt.Map(refs, func(i tag_ref.ID) xid.ID { return xid.ID(i) })
	return func(c *ent.PostMutation) {
		c.AddTagIDs(ids...)
	}
}

func WithTagsRemove(refs ...tag_ref.ID) Option {
	ids := dt.Map(refs, func(i tag_ref.ID) xid.ID { return xid.ID(i) })
	return func(c *ent.PostMutation) {
		c.RemoveTagIDs(ids...)
	}
}

type Query func(q *ent.PostQuery)

func HasKeyword(s string) Query {
	return func(q *ent.PostQuery) {
		q.Where(ent_post.Or(
			ent_post.TitleContainsFold(s),
			ent_post.SlugContainsFold(s),
			ent_post.BodyContainsFold(s),
		))
	}
}

func HasCreatedDateBefore(t time.Time) Query {
	return func(q *ent.PostQuery) {
		q.Where(ent_post.CreatedAtLT(t))
	}
}

func HasUpdatedDateBefore(t time.Time) Query {
	return func(q *ent.PostQuery) {
		q.Where(ent_post.UpdatedAtLT(t))
	}
}

func HasAuthor(id account.AccountID) Query {
	return func(q *ent.PostQuery) {
		q.Where(ent_post.HasAuthorWith(ent_account.ID(xid.ID(id))))
	}
}

func HasTags(ids []xid.ID) Query {
	return func(q *ent.PostQuery) {
		q.Where(ent_post.HasTagsWith(ent_tag.IDIn(ids...)))
	}
}

func HasCategories(ids []string) Query {
	return func(q *ent.PostQuery) {
		q.Where(ent_post.HasCategoryWith(ent_category.SlugIn(ids...)))
	}
}

func HasStatus(status ...visibility.Visibility) Query {
	pv := dt.Map(status, func(v visibility.Visibility) ent_post.Visibility { return ent_post.Visibility(v.String()) })
	return func(q *ent.PostQuery) {
		q.Where(ent_post.VisibilityIn(pv...))
	}
}

func HasNotBeenDeleted() Query {
	return func(q *ent.PostQuery) {
		q.Where(ent_post.DeletedAtIsNil())
	}
}

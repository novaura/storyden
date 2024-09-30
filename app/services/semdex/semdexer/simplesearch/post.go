package simplesearch

import (
	"context"

	"github.com/Southclaws/dt"
	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/storyden/app/resources/datagraph"
	"github.com/Southclaws/storyden/app/resources/post"
	"github.com/Southclaws/storyden/app/resources/post/post_search"
)

type postSearcher struct {
	post_search post_search.Repository
}

func (s *postSearcher) Search(ctx context.Context, query string) (datagraph.ItemList, error) {
	rs, err := s.post_search.Search(ctx, post_search.WithKeywords(query))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	items := dt.Map(rs, func(r *post.Post) datagraph.Item { return r })

	return items, nil
}

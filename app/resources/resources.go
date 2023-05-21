package resources

import (
	"go.uber.org/fx"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/authentication"
	"github.com/Southclaws/storyden/app/resources/category"
	"github.com/Southclaws/storyden/app/resources/notification"
	"github.com/Southclaws/storyden/app/resources/post"
	"github.com/Southclaws/storyden/app/resources/post_search"
	"github.com/Southclaws/storyden/app/resources/rbac"
	"github.com/Southclaws/storyden/app/resources/react"
	"github.com/Southclaws/storyden/app/resources/settings"
	"github.com/Southclaws/storyden/app/resources/tag"
	"github.com/Southclaws/storyden/app/resources/thread"
)

func Build() fx.Option {
	return fx.Options(
		rbac.Build(),
		fx.Provide(
			settings.New,
			account.New,
			authentication.New,
			category.New,
			post.New,
			tag.New,
			thread.New,
			react.New,
			notification.New,
			post_search.New,
		),
	)
}

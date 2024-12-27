package middleware

import (
	"go.uber.org/fx"

	"github.com/Southclaws/storyden/app/transports/http/middleware/chaos"
	"github.com/Southclaws/storyden/app/transports/http/middleware/frontend"
	"github.com/Southclaws/storyden/app/transports/http/middleware/limiter"
	"github.com/Southclaws/storyden/app/transports/http/middleware/origin"
	"github.com/Southclaws/storyden/app/transports/http/middleware/reqlog"
	"github.com/Southclaws/storyden/app/transports/http/middleware/session_cookie"
	"github.com/Southclaws/storyden/app/transports/http/middleware/useragent"
)

func Build() fx.Option {
	return fx.Provide(
		origin.New,
		reqlog.New,
		frontend.New,
		useragent.New,
		session_cookie.New,
		limiter.New,
		chaos.New,
	)
}

package issue

import (
	"golang.org/x/net/context"
	"github.com/opstalent/tracker/controller"
	"github.com/opstalent/tracker/router"
)

func AddRoutes(ctx context.Context) {
	router.Add("issues.list", &router.Route{
		"GET",
		"/issues",
		controller.MakeHandler(ctx, listHandler),
	})
}

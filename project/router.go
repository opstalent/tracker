package project

import (
	"golang.org/x/net/context"
	"github.com/opstalent/tracker/controller"
	"github.com/opstalent/tracker/router"
)

func AddRoutes(ctx context.Context) {
	router.Add("projects.view", &router.Route{
		"GET",
		"/projects/{id:[0-9]+}",
		controller.MakeHandler(ctx, viewHandler),
	})
}

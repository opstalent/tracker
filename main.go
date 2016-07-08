package main

import (
	"flag"
	"net/http"

	"github.com/opstalent/tracker/auth"
	"github.com/opstalent/tracker/env"
	"github.com/opstalent/tracker/issue"
	"github.com/opstalent/tracker/project"
	"github.com/opstalent/tracker/router"
	"golang.org/x/net/context"
)

var (
	username    = flag.String("u", "login", "Set redmine login, default login")
	password    = flag.String("passwd", "password", "Set redmine password, default password")
	programPort = flag.String("port", "8080", "Set server port, default 8080")
	host        = flag.String("h", "redmine.ops-dev.pl", "Set host, default = redmine.ops-dev.pl")
	port        = flag.String("p", "", "Set port, default empty")
	format      = flag.String("f", "json", "Set format, default = json")
)

func main() {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	flag.Parse()

	ctx, cancel = context.WithCancel(context.Background())
	ctx = auth.New(ctx, *username, *password, *host, *port, *format)
	defer cancel()

	issue.AddRoutes(ctx)
	project.AddRoutes(ctx)

	env.Config.Router = router.New()
	env.Config.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./resources/")))
	env.Config.Log.Critical(ctx, "%s", http.ListenAndServe(":"+*programPort, env.Config.Router))
}

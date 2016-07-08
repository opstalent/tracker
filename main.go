package main

import (
	"flag"
	"net/http"
	"strconv"

	en "github.com/caarlos0/env"
	"github.com/opstalent/tracker/auth"
	"github.com/opstalent/tracker/env"
	"github.com/opstalent/tracker/issue"
	"github.com/opstalent/tracker/project"
	"github.com/opstalent/tracker/router"
	"golang.org/x/net/context"
)

//Example program run
//PRODUCTION=true USERNAME="user" PASSWORD="password" go run tracker.go
func main() {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		host   string
		port   string
	)

	flag.StringVar(&host, "h", "redmine.ops-dev.pl", "Set host, default = redmine.ops-dev.pl")
	flag.StringVar(&port, "p", "", "Set port, default empty")
	flag.Parse()

	env.Env.Log.Critical(ctx, "%s", en.Parse(&env.Env))

	ctx, cancel = context.WithCancel(context.Background())
	ctx = auth.New(ctx, env.Env.Username, env.Env.Password, host, port, env.Env.Format)
	defer cancel()

	issue.AddRoutes(ctx)
	project.AddRoutes(ctx)

	env.Env.Router = router.New()
	env.Env.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./resources/")))
	env.Env.Log.Critical(ctx, "%s", http.ListenAndServe(":"+strconv.Itoa(env.Env.Port), env.Env.Router))
}

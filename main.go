package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/opstalent/tracker/auth"
	"github.com/opstalent/tracker/issue"
	"github.com/opstalent/tracker/logger"
	"github.com/opstalent/tracker/project"
	"github.com/opstalent/tracker/router"
	"github.com/vardius/env"
	"golang.org/x/net/context"
)

type enviroment struct {
	Username     int  `env:"USERNAME",required`
	Password     int  `env:"PASSWORD",required`
	IsProduction bool `env:"PRODUCTION"`
	Port         int  `env:"PORT" envDefault:"8080"`
	Format       int  `env:"FORMAT" envDefault:"json"`
	Log          logger.Logger
	Router       *mux.Router
}

var (
	Env = &enviroment{
		Log: logger.New(),
	}
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

	ctx, cancel = context.WithCancel(context.Background())
	ctx = auth.New(ctx, username, password, host, port, format)
	defer cancel()

	issue.AddRoutes(ctx)
	project.AddRoutes(ctx)

	Env.Router = router.New()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./resources/")))

	Env.Log.Critical(ctx, "%s", env.Parse(&Env))
	Env.Log.Critical(ctx, "%s", http.ListenAndServe(":"+strconv.Itoa(Env.Port), router))
}

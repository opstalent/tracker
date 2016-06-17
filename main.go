package main

import (
	"flag"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"github.com/opstalent/tracker/auth"
	"github.com/opstalent/tracker/controller"
	"github.com/opstalent/tracker/issue"
	"github.com/opstalent/tracker/logger"
	"github.com/opstalent/tracker/router"
)

func main() {
	var (
		ctx      context.Context
		cancel   context.CancelFunc
		username string
		password string
		host     string
		port     string
		format   string
	)

	username = os.Args[1]
	password = os.Args[2]

	flag.StringVar(&host, "f", "redmine.ops-dev.pl", "Set format, default = json")
	flag.StringVar(&port, "h", "", "Set host, default = redmine.ops-dev.pl")
	flag.StringVar(&format, "p", "json", "Set port, default empty")
	flag.Parse()

	ctx, cancel = context.WithCancel(context.Background())
	ctx = auth.NewContext(ctx, authorize(username, password, host, port, format))
	defer cancel()

	addRoutes(ctx)

	router := router.New()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./resources/")))

	log := logger.New()
	log.Critical(ctx, "%s", http.ListenAndServe(":8080", router))
}

func addRoutes(ctx context.Context) {
	router.Add("issues.list", &router.Route{
		"GET",
		"/issues",
		controller.MakeHandler(ctx, issue.ListHandler),
	})
}

func authorize(username, password, host, port, format string) auth.Auth {
	api := auth.Server{
		host,
		port,
		format,
	}
	broker := auth.Credentials{
		username,
		password,
	}

	return auth.New(auth.Api(api), auth.Broker(broker))
}

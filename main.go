package main

import (
	"flag"
	"net/http"
	"golang.org/x/net/context"

	"github.com/opstalent/tracker/auth"
	"github.com/opstalent/tracker/issue"
	"github.com/opstalent/tracker/project"
	"github.com/opstalent/tracker/logger"
	"github.com/opstalent/tracker/router"
	"os"
)

func main() {
	var (
		ctx context.Context
		cancel context.CancelFunc
		username string
		password string
		host string
		port string
		programPort string
		format string
	)

	username = os.Args[1]
	password = os.Args[2]

	flag.StringVar(&programPort, "port", "8080", "Set server port, default 8080")
	flag.StringVar(&host, "h", "redmine.ops-dev.pl", "Set host, default = redmine.ops-dev.pl")
	flag.StringVar(&port, "p", "", "Set port, default empty")
	flag.StringVar(&format, "f", "json", "Set format, default = json")
	flag.Parse()

	ctx, cancel = context.WithCancel(context.Background())
	ctx = auth.New(ctx, username, password, host, port, format)
	defer cancel()

	issue.AddRoutes(ctx)
	project.AddRoutes(ctx)

	router := router.New()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./resources/")))

	log := logger.New()
	log.Critical(ctx, "%s", http.ListenAndServe(":" + programPort, router))
}

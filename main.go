package main

import (
	"flag"
	"net/http"

	"golang.org/x/net/context"

	"github.com/opstalent/tracker/auth"
	"github.com/opstalent/tracker/issue"
	"github.com/opstalent/tracker/logger"
	"github.com/opstalent/tracker/project"
	"github.com/opstalent/tracker/router"
)

var (
	username = flag.String("u", "login", "Set redmine login, default login")
	password = flag.String("passwd", "password", "Set redmine password, default password")
	programPort = flag.String("port", "8080", "Set server port, default 8080")
	host = flag.String("h", "redmine.ops-dev.pl", "Set host, default = redmine.ops-dev.pl")
	port = flag.String("p", "", "Set port, default empty")
	format = flag.String("f", "json", "Set format, default = json")
)

func main() {
	var (
		ctx         context.Context
		cancel      context.CancelFunc
		username    string
		password    string
		host        string
		port        string
		programPort string
		format      string
	)
	flag.Parse()

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	
	ctx = auth.New(ctx, username, password, host, port, format)

	issue.AddRoutes(ctx)
	project.AddRoutes(ctx)

	router := router.New()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./resources/")))

	log := logger.New()
	log.Critical(ctx, "%s", http.ListenAndServe(":"+ programPort, router))
}

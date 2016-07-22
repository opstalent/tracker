package main

import (
	"flag"
	"net/http"

	env "github.com/opstalent/tracker/enviroment"
	"golang.org/x/net/context"
)

var programPort = flag.String("port", "8080", "Set server port, default 8080")

func main() {
	flag.Parse()
	env.Server.NotFound(http.FileServer(http.Dir("static")))
	env.Log.Critical(context.TODO(), "%s", http.ListenAndServe(":"+*programPort, env.Server))
}

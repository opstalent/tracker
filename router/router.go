package router

import (
	"github.com/gorilla/mux"
)

type (
	routerFactory func() *mux.Router
)

var (
	New routerFactory
)

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for name, route := range Routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(name).
			Handler(route.Handler)

	}

	return router
}

func init() {
	New = newRouter
}

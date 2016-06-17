package router

import (
	"net/http"
	"sync"
)

type (
	Route struct {
		Method  string
		Pattern string
		Handler http.HandlerFunc
	}
	RoutesPool map[string]*Route
)

var (
	Routes     = make(RoutesPool)
	routesLock sync.Mutex
)

func Add(name string, route *Route) {
	routesLock.Lock()
	defer routesLock.Unlock()
	Routes[name] = route
}

func Remove(name string) {
	routesLock.Lock()
	defer routesLock.Unlock()
	delete(Routes, name)
}

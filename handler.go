package main

import (
	"html/template"
	"net/http"
	"time"

	env "github.com/opstalent/tracker/enviroment"
	"github.com/vardius/goserver"
	"golang.org/x/net/context"
)

type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request, *goserver.Context)

func NewHandler(h HandlerFunc) goserver.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c *goserver.Context) {
		start := time.Now()
		ctx, cancel, err := newContext(r)
		if err != nil {
			panic(err)
		}
		defer cancel()
		h(ctx, w, r, c)
		env.Log.Info(ctx, "%s\t%s\t%d", r.Method, r.RequestURI, time.Since(start))
	}
}

func Render(w http.ResponseWriter, tmpl *template.Template, args interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := tmpl.Execute(w, args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

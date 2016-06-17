package controller

import (
	"net/http"
	"time"

	"github.com/opstalent/tracker/logger"
	"golang.org/x/net/context"
)

type (
	handlerFnc func(ctx context.Context, w http.ResponseWriter, r *http.Request)
)

func MakeHandler(ctx context.Context, fn handlerFnc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log := logger.New()

		log.Info(ctx, "%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start))

		fn(ctx, w, r)
	}
}

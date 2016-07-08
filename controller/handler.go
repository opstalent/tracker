package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/opstalent/tracker/env"
	"golang.org/x/net/context"
)

type (
	handlerFnc func(ctx context.Context, w http.ResponseWriter, r *http.Request)
)

func MakeHandler(ctx context.Context, fn handlerFnc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		env.Config.Log.Info(ctx, "%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start))

		fn(ctx, w, r)
	}
}

func CallAPI(ctx context.Context, r *http.Request, url string, entity interface{}) error {
	var (
		err error
		req *http.Request
	)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	for k, vs := range r.URL.Query() {
		q.Set(k, vs[0])
	}

	req.URL.RawQuery = q.Encode()

	return httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err = json.NewDecoder(resp.Body).Decode(entity); err != nil {
			return err
		}

		return nil
	})
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() {
		c <- f(client.Do(req))
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}

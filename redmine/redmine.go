package redmine

import (
	"encoding/json"
	"flag"
	"net/http"

	"golang.org/x/net/context"
)

var (
	Username = flag.String("u", "login", "Set redmine login, default login")
	Password = flag.String("passwd", "password", "Set redmine password, default password")
	Host     = flag.String("h", "redmine.ops-dev.pl", "Set host, default = redmine.ops-dev.pl")
	Port     = flag.String("p", "", "Set port, default empty")
	Format   = flag.String("f", "json", "Set format, default = json")
)

func GetURL(params ...string) string {
	if len(params) < 1 {
		panic("redmine: redmine.GetUrl method, not enought params!")
	}
	var portSeparator string
	if *Port != "" {
		portSeparator = ":"
	}
	var paramsStr string
	for i := 0; i < len(params); i++ {
		paramsStr += "/" + params[i]
	}
	return "http://" + *Username + ":" + *Password + "@" + *Host + portSeparator + *Port + paramsStr + ".json"
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

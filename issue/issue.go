package issue

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/opstalent/tracker/auth"
	"github.com/opstalent/tracker/logger"
	"github.com/opstalent/tracker/resource"
	"golang.org/x/net/context"
)

const (
	prefix = "issues"
)

type (
	Issues struct {
		Resources []Issue `json:"issues"`
		Total     int     `json:"total_count"`
		Limit     int     `json:"limit"`
		Offset    int     `json:"offset"`
	}

	Issue struct {
		Id           int                    `json:"id"`
		DoneRatio    int                    `json:"done_ratio"`
		Subject      string                 `json:"subject"`
		Description  string                 `json:"description"`
		StartDate    string                 `json:"start_date"`
		CreatedOn    time.Time              `json:"created_on"`
		UpdatedOn    time.Time              `json:"updated_on"`
		Project      resource.Field         `json:"project"`
		Tracker      resource.Field         `json:"tracker"`
		Status       resource.Field         `json:"status"`
		Priority     resource.Field         `json:"priority"`
		Author       resource.Field         `json:"author"`
		AssignedTo   resource.Field         `json:"assigned_to"`
		CustomFields []resource.CustomField `json:"custom_fields"`
	}
)

func Get(ctx context.Context, v url.Values) (*Issues, error) {
	var (
		err    error
		url    string
		req    *http.Request
		issues = &Issues{}
	)

	if a, ok := auth.FromContext(ctx); ok {
		opts := a.Options()
		url = "http://" + opts.Broker.User + ":" + opts.Broker.Password + "@" + opts.Api.Host + opts.Api.Port + "/" + prefix + ".json"
	}

	log := logger.New()
	log.Info(ctx, "%s", url)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, vs := range v {
		q.Set(k, vs[0])
	}

	req.URL.RawQuery = q.Encode()

	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err = json.NewDecoder(resp.Body).Decode(issues); err != nil {
			return err
		}

		return nil
	})

	return issues, err
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}

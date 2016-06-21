package issue

import (
	"net/http"
	"github.com/opstalent/tracker/auth"
	"golang.org/x/net/context"
	"github.com/opstalent/tracker/controller"
)

func Get(ctx context.Context, r *http.Request) (*Issues, error) {
	var (
		err error
		url string
		issues = &Issues{}
	)

	if a, ok := auth.FromContext(ctx); ok {
		opts := a.Options()
		url = "http://" + opts.Broker.User + ":" + opts.Broker.Password + "@" + opts.Api.Host + opts.Api.Port + "/" + prefix + ".json"
	}

	err = controller.CallAPI(ctx, r, url, issues)

	return issues, err
}

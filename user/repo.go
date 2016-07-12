package user

import (
	"net/http"

	"github.com/opstalent/tracker/auth"
	"github.com/opstalent/tracker/controller"
	"golang.org/x/net/context"
)

func Get(ctx context.Context, r *http.Request) (*Users, error) {
	var (
		err   error
		url   string
		users = &Users{}
	)

	if a, ok := auth.FromContext(ctx); ok {
		opts := a.Options()
		url = "http://" + opts.Broker.User + ":" + opts.Broker.Password + "@" + opts.Api.Host + opts.Api.Port + "/" + prefix + ".json"
	}

	err = controller.CallAPI(ctx, r, url, users)

	return users, err
}

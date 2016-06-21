package project

import (
	"net/http"
	"github.com/opstalent/tracker/auth"
	"golang.org/x/net/context"
	"github.com/opstalent/tracker/controller"
)

func getById(ctx context.Context, r *http.Request, id string) (*Project, error) {
	var (
		err error
		url string
		project = struct{ Project Project `json:"project"` }{}
	)

	if a, ok := auth.FromContext(ctx); ok {
		opts := a.Options()
		url = "http://" + opts.Broker.User + ":" + opts.Broker.Password + "@" + opts.Api.Host + opts.Api.Port + "/" + prefix + "/" + id + ".json"
	}

	err = controller.CallAPI(ctx, r, url, &project)

	return &project.Project, err
}

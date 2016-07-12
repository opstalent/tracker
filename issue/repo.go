package issue

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/opstalent/tracker/auth"
	"github.com/opstalent/tracker/controller"
	"github.com/opstalent/tracker/user"
	"golang.org/x/net/context"
)

func Get(ctx context.Context, r *http.Request, user *user.User) (*Issues, error) {
	var (
		err    error
		url    string
		issues = &Issues{}
	)

	if a, ok := auth.FromContext(ctx); ok {
		opts := a.Options()
		url = "http://" +
			opts.Broker.User +
			":" + opts.Broker.Password +
			"@" + opts.Api.Host + opts.Api.Port +
			"/" + prefix +
			".json"

		if user.Id > 0 {
			url += "?assigned_to_id=" + strconv.Itoa(user.Id)
		}
	}
	fmt.Println(url)

	err = controller.CallAPI(ctx, r, url, issues)

	return issues, err
}

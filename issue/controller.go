package issue

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/opstalent/tracker/env"
	"github.com/opstalent/tracker/user"
	"golang.org/x/net/context"
)

var (
	funcs = template.FuncMap{"countTotal": countTotal}
	tmpl  = template.Must(template.New("list.html").Funcs(funcs).ParseFiles("views/issue/list.html"))
)

type Data struct {
	List SortedIssues
}

func listHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var issues = &Issues{}
	users, err := user.Get(ctx, r)
	fmt.Println("USERS")
	fmt.Println(users)
	for _, user := range users.Resources {
		if is, err := Get(ctx, r, &user); err == nil {
			fmt.Println("LOOP")
			fmt.Println(is)
			issues.Resources = append(issues.Resources, is.Resources...)
		}
	}
	fmt.Println("ISSUES")
	fmt.Println(issues)
	if err != nil {
		env.Config.Log.Critical(ctx, "%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		data := &Data{Sort(issues)}
		render(w, data)
	}
}

func render(w http.ResponseWriter, args *Data) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := tmpl.Execute(w, args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func countTotal(s map[string][]Issue) int {
	total := 0
	for _, p := range s {
		total += len(p)
	}

	return total
}

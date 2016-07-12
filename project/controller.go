package project

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/opstalent/tracker/env"
	"github.com/opstalent/tracker/issue"
	"github.com/opstalent/tracker/resource"
	"github.com/opstalent/tracker/user"
	"golang.org/x/net/context"
)

var (
	funcs = template.FuncMap{"getUsers": getUsers}
	tmpl  = template.Must(template.New("view.html").Funcs(funcs).ParseFiles("views/project/view.html"))
)

func viewHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project, err := GetById(ctx, r, vars["id"])
	if err != nil {
		env.Config.Log.Critical(ctx, "%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		req, err := http.NewRequest("GET", "http://notimportant.com", nil)
		if err != nil {
			env.Config.Log.Critical(ctx, "%s", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		q := req.URL.Query()
		q.Set("project_id", strconv.Itoa(project.Id))
		q.Set("limit", strconv.Itoa(9999))
		req.URL.RawQuery = q.Encode()
		is, err := issue.Get(ctx, req, &user.User{})

		project.Issues = issue.SortByStatus(is)

		render(w, project)
	}
}

func render(w http.ResponseWriter, args *Project) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := tmpl.Execute(w, args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getUsers(is []issue.Issue) map[string]resource.Field {
	users := make(map[string]resource.Field)
	for _, issue := range is {
		_, ok := users[issue.AssignedTo.Name]
		if len(issue.AssignedTo.Name) != 0 && !ok {
			users[issue.AssignedTo.Name] = issue.AssignedTo
		}
	}
	return users
}

package project

import (
	"html/template"
	"net/http"

	"github.com/opstalent/tracker/logger"
	"golang.org/x/net/context"
	"github.com/gorilla/mux"
	"github.com/opstalent/tracker/issue"
	"strconv"
)

var (
	tmpl = template.Must(template.New("view.html").ParseFiles("views/project/view.html"))
)

func viewHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project, err := GetById(ctx, r, vars["id"])
	log := logger.New()
	if err != nil {
		log.Critical(ctx, "%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		req, err := http.NewRequest("GET", "http://notimportant.com", nil)
		if err != nil {
			log.Critical(ctx, "%s", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		q := req.URL.Query()
		q.Set("project_id", strconv.Itoa(project.Id))
		q.Set("limit", strconv.Itoa(9999))
		req.URL.RawQuery = q.Encode()
		is, err := issue.Get(ctx, req)

		project.Issues = is

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

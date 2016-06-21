package project

import (
	"html/template"
	"net/http"

	"github.com/opstalent/tracker/logger"
	"golang.org/x/net/context"
	"github.com/gorilla/mux"
)

var (
	tmpl = template.Must(template.New("view.html").ParseFiles("views/project/view.html"))
)

func viewHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project, err := getById(ctx, r, vars["id"])
	if err != nil {
		log := logger.New()
		log.Critical(ctx, "%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
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

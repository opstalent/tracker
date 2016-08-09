package main

import (
	"net/http"
	"strconv"

	env "github.com/opstalent/tracker/enviroment"
	"github.com/vardius/goserver"
	"golang.org/x/net/context"
)

func init() {
	env.Server.GET("/projects/:id:[0-9]+", NewHandler(projectViewHandler))
	env.Server.GET("/issues", NewHandler(issueListHandler))
}

func issueListHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, c *goserver.Context) {
	var users = new(Users)
	if err := users.Get(ctx, r); err != nil {
		env.Log.Critical(ctx, "%s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	Render(w, issueListTemplate, users)
}

func projectViewHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, c *goserver.Context) {
	var project = new(Project)
	if err := project.Get(ctx, r, c.Params["id"]); err != nil {
		env.Log.Critical(ctx, "%s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest("GET", "http://notimportant.com", nil)
	if err != nil {
		env.Log.Critical(ctx, "%s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	q := req.URL.Query()
	q.Set("project_id", strconv.Itoa(project.Id))
	q.Set("offset", strconv.Itoa(0))
	q.Set("limit", strconv.Itoa(9999))
	req.URL.RawQuery = q.Encode()

	var issues = new(Issues)
	if err := issues.Get(ctx, req, nil); err != nil {
		env.Log.Critical(ctx, "%s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	Render(w, projectViewTemplate, struct {
		Project *Project
		Issues  *Issues
	}{project, issues})
}

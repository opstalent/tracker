package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/opstalent/tracker/redmine"
	"golang.org/x/net/context"
)

type (
	Projects struct {
		Resources []*Project `json:"projects"`
		Total     int        `json:"total_count"`
		Limit     int        `json:"limit"`
		Offset    int        `json:"offset"`
	}

	Project struct {
		Id          int       `json:"id"`
		Identifier  string    `json:"identifier"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Homepage    string    `json:"homepage"`
		Status      int       `json:"status"`
		Parent      Field     `json:"parent"`
		Created     time.Time `json:"created_on"`
		Updated     time.Time `json:"updated_on"`
	}
)

func (projects *Projects) Get(ctx context.Context, r *http.Request, user *User) error {
	url := redmine.GetUrl("projects")
	if user.Id > 0 {
		url += "?assigned_to_id=" + strconv.Itoa(user.Id)
	}
	return redmine.CallAPI(ctx, r, url, projects)
}

func (project *Project) Get(ctx context.Context, r *http.Request, id string) error {
	var data struct {
		Project *Project `json:"project"`
	}
	data.Project = project
	url := redmine.GetUrl("projects", id)
	return redmine.CallAPI(ctx, r, url, &data)
}

package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	env "github.com/opstalent/tracker/enviroment"
	"golang.org/x/net/context"
)

var (
	projectView = template.Must(template.New("view.html").Funcs(template.FuncMap{
		"getUsers":        getUsers,
		"statusIdOfFirst": statusIdOfFirst,
	}).ParseFiles("views/projects/view.html"))
	issueList = template.Must(template.New("list.html").Funcs(template.FuncMap{
		"getDonut":            getDonut,
		"getUserIssues":       getUserIssues,
		"getIssuesPerProject": getIssuesPerProject,
	}).ParseFiles("views/issues/list.html"))
)

func getUsers(is []*Issue) map[string]int {
	users := make(map[string]int)
	for _, issue := range is {
		if len(issue.AssignedTo.Name) != 0 {
			users[issue.AssignedTo.Name] += 1
		}
	}
	return users
}

func statusIdOfFirst(is []*Issue) int {
	for _, issue := range is {
		return issue.Status.Id
	}
	return 0
}

func getUserIssues(user *User) *Issues {
	ctx := context.TODO()
	req, err := http.NewRequest("GET", "http://notimportant.com", nil)
	if err != nil {
		env.Log.Critical(ctx, "%s", err)
		return nil
	}
	q := req.URL.Query()
	q.Set("offset", strconv.Itoa(0))
	q.Set("limit", strconv.Itoa(9999))
	req.URL.RawQuery = q.Encode()

	var issues = new(Issues)
	if err := issues.Get(ctx, req, user); err != nil {
		env.Log.Critical(ctx, "%s", err)
		return nil
	}
	return issues
}

func getIssuesPerProject(issues []*Issue) SortedIssues {
	sorted := make(SortedIssues)
	for _, issue := range issues {
		sorted[issue.Project.Name] = append(sorted[issue.Project.Name], issue)
	}
	return sorted
}

func getDonut(issues SortedIssues) template.JS {
	type element struct {
		Label string `json:"label"`
		Value int    `json:"value"`
	}
	var data []element
	for status, issues := range issues {
		data = append(data, element{status, len(issues)})
	}
	b, err := json.Marshal(data)
	if err != nil {
		env.Log.Info(ctx, "%s", err)
		return ""
	}
	return template.JS(b)
}

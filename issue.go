package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/opstalent/tracker/redmine"
	"golang.org/x/net/context"
)

type (
	Issues struct {
		Resources []*Issue `json:"issues"`
		Total     int      `json:"total_count"`
		Limit     int      `json:"limit"`
		Offset    int      `json:"offset"`
	}

	Issue struct {
		Id           int           `json:"id"`
		DoneRatio    int           `json:"done_ratio"`
		Subject      string        `json:"subject"`
		Description  string        `json:"description"`
		StartDate    string        `json:"start_date"`
		Created      time.Time     `json:"created_on"`
		Updated      time.Time     `json:"updated_on"`
		Project      Field         `json:"project"`
		Tracker      Field         `json:"tracker"`
		Status       Field         `json:"status"`
		Priority     Field         `json:"priority"`
		Author       Field         `json:"author"`
		AssignedTo   Field         `json:"assigned_to"`
		CustomFields []CustomField `json:"custom_fields"`
	}
	SortedIssues map[string][]*Issue
)

func (issues *Issues) Get(ctx context.Context, r *http.Request, user *User) error {
	url := redmine.GetURL("issues")
	if user != nil && user.Id > 0 {
		url += "?assigned_to_id=" + strconv.Itoa(user.Id)
	}
	return redmine.CallAPI(ctx, r, url, issues)
}

func (issues *Issues) SortByStatus() SortedIssues {
	sorted := make(SortedIssues)
	for _, issue := range issues.Resources {
		sorted[issue.Status.Name] = append(sorted[issue.Status.Name], issue)
	}

	return sorted
}

func (issues *Issues) SortByProject() SortedIssues {
	sorted := make(SortedIssues)
	for _, issue := range issues.Resources {
		sorted[issue.Project.Name] = append(sorted[issue.Project.Name], issue)
	}
	return sorted
}

func (issues *Issues) SortByTracker() SortedIssues {
	sorted := make(SortedIssues)
	for _, issue := range issues.Resources {
		sorted[issue.Tracker.Name] = append(sorted[issue.Tracker.Name], issue)
	}
	return sorted
}

func (issues *Issues) SortByPriority() SortedIssues {
	sorted := make(SortedIssues)
	for _, issue := range issues.Resources {
		sorted[issue.Priority.Name] = append(sorted[issue.Priority.Name], issue)
	}
	return sorted
}

func (issues *Issues) SortByAuthor() SortedIssues {
	sorted := make(SortedIssues)
	for _, issue := range issues.Resources {
		sorted[issue.Author.Name] = append(sorted[issue.Author.Name], issue)
	}
	return sorted
}

func (issues *Issues) SortByAssigned() SortedIssues {
	sorted := make(SortedIssues)
	for _, issue := range issues.Resources {
		sorted[issue.AssignedTo.Name] = append(sorted[issue.AssignedTo.Name], issue)
	}
	return sorted
}

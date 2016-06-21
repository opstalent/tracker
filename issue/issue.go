package issue

import (
	"time"

	"github.com/opstalent/tracker/resource"
)

const (
	prefix = "issues"
)

type (
	Issues struct {
		Resources []Issue `json:"issues"`
		Total     int     `json:"total_count"`
		Limit     int     `json:"limit"`
		Offset    int     `json:"offset"`
	}

	Issue struct {
		Id           int                    `json:"id"`
		DoneRatio    int                    `json:"done_ratio"`
		Subject      string                 `json:"subject"`
		Description  string                 `json:"description"`
		StartDate    string                 `json:"start_date"`
		Created      time.Time              `json:"created_on"`
		Updated      time.Time              `json:"updated_on"`
		Project      resource.Field         `json:"project"`
		Tracker      resource.Field         `json:"tracker"`
		Status       resource.Field         `json:"status"`
		Priority     resource.Field         `json:"priority"`
		Author       resource.Field         `json:"author"`
		AssignedTo   resource.Field         `json:"assigned_to"`
		CustomFields []resource.CustomField `json:"custom_fields"`
	}
)

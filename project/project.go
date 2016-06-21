package project

import (
	"time"
	"github.com/opstalent/tracker/resource"
)

const (
	prefix = "projects"
)

type (
	Project struct {
		Id          int                      `json:"id"`
		identifier  string                   `json:"identifier"`
		Name        string                   `json:"name"`
		Description string                   `json:"description"`
		Homepage    string                   `json:"homepage"`
		Status      int                      `json:"status"`
		//CustomFields []resource.CustomField   `json:"custom_fields"`
		Parent      resource.Field           `json:"parent"`
		Created     time.Time                `json:"created_on"`
		Updated     time.Time                `json:"updated_on"`
	}
)

package main

import (
	"net/http"
	"time"

	"github.com/opstalent/tracker/redmine"
	"golang.org/x/net/context"
)

type (
	Users struct {
		Resources []*User `json:"users"`
		Total     int     `json:"total_count"`
		Limit     int     `json:"limit"`
		Offset    int     `json:"offset"`
	}

	User struct {
		Id        int       `json:"id"`
		Login     string    `json:"login"`
		FirstName string    `json:"firstname"`
		LastName  string    `json:"lastname"`
		Email     string    `json:"mail"`
		Created   time.Time `json:"created_on"`
		LastLogin time.Time `json:"last_login_on"`
	}
)

func (users *Users) Get(ctx context.Context, r *http.Request) error {
	url := redmine.GetUrl("users")
	return redmine.CallAPI(ctx, r, url, users)
}

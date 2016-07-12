package user

import "time"

const (
	prefix = "users"
)

type (
	Users struct {
		Resources []User `json:"users"`
		Total     int    `json:"total_count"`
		Limit     int    `json:"limit"`
		Offset    int    `json:"offset"`
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

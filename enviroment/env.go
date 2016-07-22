package env

import (
	"github.com/vardius/goapi"
	"github.com/vardius/golog"
)

var (
	Log    = golog.New()
	Server = goapi.New()
)

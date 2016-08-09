package env

import (
	"github.com/vardius/golog"
	"github.com/vardius/goserver"
)

var (
	Log    = golog.New()
	Server = goserver.New()
)

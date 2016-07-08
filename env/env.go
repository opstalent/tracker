package env

import (
	"github.com/gorilla/mux"
	"github.com/opstalent/tracker/logger"
)

type enviroment struct {
	Log    logger.Logger
	Router *mux.Router
}

var (
	Config = &enviroment{
		Log: logger.New(),
	}
)

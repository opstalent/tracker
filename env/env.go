package env

import (
	"github.com/gorilla/mux"
	"github.com/opstalent/tracker/logger"
)

type enviroment struct {
	Username     string `env:"USERNAME",required`
	Password     string `env:"PASSWORD",required`
	IsProduction bool   `env:"PRODUCTION"`
	Port         int    `env:"PORT" envDefault:"8080"`
	Format       string `env:"FORMAT" envDefault:"json"`
	Log          logger.Logger
	Router       *mux.Router
}

var (
	Env = &enviroment{
		Log: logger.New(),
	}
)

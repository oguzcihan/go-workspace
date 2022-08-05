package app

import (
	"github.com/core-go/core/log"
	"github.com/core-go/core/middleware"
	"github.com/core-go/core/sql"
	sv "github.com/core-go/service"
)

type Config struct {
	Server     sv.ServerConf        `mapstructure:"server"`
	Sql        sql.Config           `mapstructure:"sql"`
	Log        log.Config           `mapstructure:"log"`
	MiddleWare middleware.LogConfig `mapstructure:"middleware"`
}

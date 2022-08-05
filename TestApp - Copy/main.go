package main

import (
	"TestApp/internal/app"
	"context"
	"fmt"
	"github.com/core-go/core/config"
	"github.com/core-go/core/log"
	"github.com/core-go/core/middleware"
	. "github.com/core-go/service"
	"github.com/gorilla/mux"
)

func main() {

	var conf app.Config
	err := config.Load(&conf, "configs/config")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	log.Initialize(conf.Log)
	r.Use(middleware.BuildContext)
	logger := middleware.NewLogger()
	if log.IsInfoEnable() {
		r.Use(middleware.Logger(conf.MiddleWare, log.InfoFields, logger))
	}
	r.Use(middleware.Recover(log.PanicMsg))
	err = app.Route(r, context.Background(), conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(ServerInfo(conf.Server))
	server := CreateServer(conf.Server, r)
	if err = server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}

}

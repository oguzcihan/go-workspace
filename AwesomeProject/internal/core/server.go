package core

import (
	helper "AwesomeProject/internal/helpers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

var (
	port   = ":9090"
	router = mux.NewRouter()
)

func StartServer() {
	helper.InitializeLogger()

	helper.Logger.Info("---connecting routes---")
	allRoutes()
	helper.Logger.Info("---connected---")

	errRunServe := http.ListenAndServe(port, router)
	if errRunServe != nil {
		helper.Logger.Fatal("router_run_error:", zap.Error(errRunServe))
	}
}

func allRoutes() {
	UserRoute(router)
	AccountRoute(router)

}

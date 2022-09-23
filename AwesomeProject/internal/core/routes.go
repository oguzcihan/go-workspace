package core

import (
	helper "AwesomeProject/internal/helpers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var (
	application, err = NewApplication()
)

func UserRoute(router *mux.Router) {
	if err != nil {
		helper.Logger.Fatal("application_error", zap.Error(err))

	}
	router.HandleFunc("/user", application.UserHandler.Create).Methods("POST")
	router.HandleFunc("/user/{id}", application.UserHandler.Save).Methods("PUT")
	router.HandleFunc("/user/{id}", application.UserHandler.Delete).Methods("DELETE")
	router.HandleFunc("/user", application.UserHandler.GetAll).Methods("GET")

}

func AccountRoute(router *mux.Router) {

	router.HandleFunc("/register", application.AccountHandler.Register).Methods("POST")
	router.HandleFunc("/login", application.AccountHandler.Login).Methods("POST")
	//
	//router.HandleFunc("/user", IsAuthorized(UserIndex)).Methods("GET")
	//router.HandleFunc("/admin", IsAuthorized(AdminIndex)).Methods("GET")
	//router.HandleFunc("/", Index).Methods("GET")
	//router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Access-Control-Allow-Origin", "")
	//	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	//})
}

package core

import (
	. "AwesomeProject/internal/handlers"
	. "AwesomeProject/internal/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func UserRoute(userHandler IUserHandler, router *mux.Router) {
	//router.HandleFunc("/user", IsAuthorized(userHandler.Create)).Methods("POST")
	router.HandleFunc("/user/{id}", IsAuthorized(userHandler.Save)).Methods("PUT")
	router.HandleFunc("/user/{id}", IsAuthorized(userHandler.Delete)).Methods("DELETE")
	router.HandleFunc("/users", IsAuthorized(userHandler.GetAll)).Methods("GET")

}

func AccountRoute(accountHandler IAccountHandler, router *mux.Router) {
	router.HandleFunc("/signup", accountHandler.Register).Methods("POST")
	router.HandleFunc("/login", accountHandler.Login).Methods("POST")

	router.HandleFunc("/user", IsAuthorized(accountHandler.UserIndex)).Methods("GET")
	router.HandleFunc("/admin", IsAuthorized(accountHandler.AdminIndex)).Methods("GET")
	router.HandleFunc("/", accountHandler.Index).Methods("GET")

	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
}

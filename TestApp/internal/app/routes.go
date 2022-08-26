package app

import (
	. "context"
	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router, context Context) error {
	application, err := NewApplication(context)
	if err != nil {
		return err
	}
	user := "/user"
	r.HandleFunc(user, application.User.Create).Methods("POST")
	r.HandleFunc(user+"/{id}", application.User.Save).Methods("PUT")
	r.HandleFunc(user+"/{id}", application.User.Delete).Methods("DELETE")
	r.HandleFunc(user, application.User.GetAllUser).Methods("GET")

	return nil
}

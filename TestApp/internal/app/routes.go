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
	r.HandleFunc(user+"/{id}", application.User.Update).Methods("PUT")

	return nil
}

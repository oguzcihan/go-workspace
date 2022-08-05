package app

import (
	"context"
	. "github.com/core-go/core"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}
	user := "/users"
	r.HandleFunc(user, app.User.Create).Methods(POST)
	r.HandleFunc(user+"/{id}", app.User.Update).Methods(PUT)
	r.HandleFunc(user+"/{id}", app.User.Delete).Methods(DELETE)
	return nil
}

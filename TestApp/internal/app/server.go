package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RunServer() {
	r := mux.NewRouter()
	fmt.Println("connecting to routes")

	err := UserRoute(r, context.Background())
	if err != nil {
		return
	}
	fmt.Println("connected")

	log.Fatal(http.ListenAndServe(":9090", r))
}

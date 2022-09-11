package app

import (
	. "TestApp/internal/utils"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RunServer() {
	InitializeLogger()
	r := mux.NewRouter()

	Logger.Info("Connecting routes")
	err := UserRoute(r, context.Background())
	if err != nil {
		Logger.Error("routes_error")
		return
	}
	Logger.Info("Connected")

	log.Fatal(http.ListenAndServe(":9090", r))
}

/*
	context.Background(): Boş bir contexttir, deadline ı bulunmaz iptal edilemez
	üretilen contextler için taban olarak kullanılabilir
	context.TODO(): Ne kullanılacağından emin olmadığında kullanılır
*/

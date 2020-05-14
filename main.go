package main

import (
	"github.com/dorokhin/sample-api/handlers"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Get("/api/users", handlers.GetUsers)
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

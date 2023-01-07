package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-queue-management-http/configs"
	"go-queue-management-http/handlers"
	"net/http"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)
	r.Get("/", handlers.List)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}

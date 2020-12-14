package main

import (
	"net/http"

	"github.com/buffup/GolangTechTask/cmd/server/internal/handlers"
)

func main() {
	// Replace with DB impl
	store := handlers.NewInMemStore()

	routes := handlers.Routes(store)
	if err := http.ListenAndServe(":8080", routes); err != nil {
		panic(err)
	}
}

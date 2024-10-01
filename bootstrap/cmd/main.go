package main

import (
	"COREMOD/app/router"

	"github.com/go-chi/chi/v5"
)

func main() {
	chiRouter := chi.NewMux()

	router.InitializeMiddleware(chiRouter)
	router.InitializeRoutes(chiRouter)
}

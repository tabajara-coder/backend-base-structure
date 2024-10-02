package main

import (
	"COREMOD/app"
	"net/http"

	"github.com/tabajara-coder/backend-base-structure/core"

	"github.com/go-chi/chi/v5"
)

func main() {
	chiRouter := chi.NewMux()

	app.InitializeMiddleware(chiRouter)
	app.InitializeRoutes(chiRouter)

	listenAddr := core.Getenv("HTTP_LISTEN_ADDR", ":8080")

	http.ListenAndServe(listenAddr, chiRouter)
}

package main

import (
	"COREMOD/infra/router"

	"github.com/tabajara-coder/common/env"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	chiRouter := chi.NewMux()

	router.InitializeMiddleware(chiRouter)
	router.InitializeRoutes(chiRouter)

	listenAddr := env.Get("HTTP_LISTEN_ADDR", ":8080")

	http.ListenAndServe(listenAddr, chiRouter)
}

func init() {
	env.Setup()
}

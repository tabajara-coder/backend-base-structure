package main

import (
	"COREMOD/app"
	"COREMOD/service/env"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	chiRouter := chi.NewMux()

	app.InitializeMiddleware(chiRouter)
	app.InitializeRoutes(chiRouter)

	listenAddr := env.Get("HTTP_LISTEN_ADDR", ":8080")

	http.ListenAndServe(listenAddr, chiRouter)
}
func init() {
	env.Setup()
}

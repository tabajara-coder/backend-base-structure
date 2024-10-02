package main

import (
	"COREMOD/app"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tabajara-coder/backend-base-structure/core"

	"github.com/go-chi/chi/v5"
)

func main() {
	chiRouter := chi.NewMux()

	core.Setup()

	app.InitializeMiddleware(chiRouter)
	app.InitializeRoutes(chiRouter)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDR")

	http.ListenAndServe(listenAddr, chiRouter)
}
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

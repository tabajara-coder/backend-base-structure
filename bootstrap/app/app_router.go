package app

import (
	"COREMOD/app/service/auth"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/tabajara-coder/backend-base-structure/core"
)

func InitializeMiddleware(router *chi.Mux) {
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
}

func InitializeRoutes(router *chi.Mux) {

	auth.InitializeAuthRouter(router)

	router.Get("/", core.Handler(HandleSign))
}

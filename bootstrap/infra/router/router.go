package router

import (
	"COREMOD/adapter/driving/auth"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/tabajara-coder/common/core"
)

func InitializeMiddleware(router *chi.Mux) {
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
}

func InitializeRoutes(router *chi.Mux) {

	auth.InitializeAuthRouter(router)

	router.Get("/", core.Handler(HandleSign))
}

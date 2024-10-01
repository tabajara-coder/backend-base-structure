package router

import (
	"COREMOD/app/handler"

	"github.com/go-chi/chi/v5"
	"github.com/tabajara-coder/backend-base-structure/core"
)

func InitializeAuthRouter(router chi.Router) {
	authConfig := core.AuthenticationConfig{
		AuthFunc: handler.VerifyUserSession,
	}

	router.Use(core.WithAuthentication(authConfig))
}

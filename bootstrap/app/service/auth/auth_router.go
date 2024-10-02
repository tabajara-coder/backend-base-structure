package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/tabajara-coder/backend-base-structure/core"
)

func InitializeAuthRouter(router chi.Router) {
	authConfig := core.AuthenticationConfig{
		AuthFunc: VerifyUserSession,
	}

	router.Use(core.WithAuthentication(authConfig))
}

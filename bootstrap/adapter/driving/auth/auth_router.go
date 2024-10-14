package auth

import (
	"COREMOD/domain/service/auth"

	"github.com/go-chi/chi/v5"
	"github.com/tabajara-coder/common/core"
)

func InitializeAuthRouter(router chi.Router) {
	authConfig := core.AuthenticationConfig{
		AuthFunc: auth.VerifyUserSession,
	}

	router.Use(core.WithAuthentication(authConfig))
}

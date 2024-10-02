package auth

import (
	"github.com/tabajara-coder/backend-base-structure/core"
)

func VerifyUserSession(core *core.Core) (core.Auth, error) {

	//verifica no banco a sessão
	return Auth{
		LoggedIn: true,
		UserID:   2333,
		Email:    "",
	}, nil
}

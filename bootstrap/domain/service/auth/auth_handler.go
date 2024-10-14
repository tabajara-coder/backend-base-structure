package auth

import (
	"COREMOD/domain/entity"

	"github.com/tabajara-coder/common/core"
)

func VerifyUserSession(core *core.Core) (core.Auth, error) {

	//verifica no banco a sessão
	return entity.Auth{
		LoggedIn: true,
		UserID:   2333,
		Email:    "",
	}, nil
}

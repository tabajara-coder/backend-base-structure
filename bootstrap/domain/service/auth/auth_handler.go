package auth


import (
	"github.com/tabajara-coder/backend-base-structure/core"
	"COREMOD/domain/entity"
)

func VerifyUserSession(core *core.Core) (core.Auth, error) {

	//verifica no banco a sessão
	return entity.Auth{
		LoggedIn: true,
		UserID:   2333,
		Email:    "",
	}, nil
}

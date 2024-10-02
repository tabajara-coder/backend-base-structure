package app

import (
	"net/http"

	"github.com/tabajara-coder/backend-base-structure/core"
)

func HandleSign(core *core.Core) error {
	return core.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

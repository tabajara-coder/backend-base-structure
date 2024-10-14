package router

import (
	"net/http"

	"github.com/tabajara-coder/common/core"
)

func HandleSign(core *core.Core) error {
	return core.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

package core

import (
	"encoding/json"
	"net/http"
	"os"
)

type Core struct {
	Response http.ResponseWriter
	Request  *http.Request
}

type HandlerFunc func(core *Core) error

type Auth interface {
	Check() bool
}

func JSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func Handler(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		core := &Core{
			Response: w,
			Request:  r,
		}
		if err := h(core); err != nil {
			JSON(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func (core *Core) Getenv(name string, def string) string {
	return Getenv(name, def)
}

func Getenv(name string, def string) string {
	env := os.Getenv(name)
	if len(env) == 0 {
		return def
	}
	return env
}

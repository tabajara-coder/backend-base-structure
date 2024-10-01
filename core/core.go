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

func (core *Core) JSON(status int, v any) error {
	return JSON(core.Response, status, v)
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

type AuthenticationConfig struct {
	AuthFunc func(*Core) (Auth, error)
}

func WithAuthentication(config AuthenticationConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			kit := &Core{
				Response: w,
				Request:  r,
			}
			auth, err := config.AuthFunc(kit)
			if err != nil {
				JSON(w, http.StatusInternalServerError, err.Error())
				return
			}
			if !auth.Check() {
				JSON(w, http.StatusForbidden, nil)
			}
			next.ServeHTTP(w, r)
		})
	}
}

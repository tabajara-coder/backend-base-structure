package env

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type envService struct{}

var (
	envInstance *envService
	once        sync.Once
)

func Env() *envService {
	once.Do(func() {
		envInstance = &envService{}
	})
	return envInstance
}

func (env *envService) get(name string, def string) string {
	val := os.Getenv(name)
	if val == "" {
		return def
	}
	return val
}

func (env *envService) setup() {
	if err := godotenv.Load("bootstrap/.env"); err != nil {
		log.Fatal(err)
	}
}

func Get(name string, def string) string {
	return Env().get(name, def)
}

func Setup() {
	Env().setup()
}

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DriverPostgres = "postgres"
	DriverSqlite3  = "sqlite3"
	DriverMysql    = "mysql"
)

type Config struct {
	Driver   string
	Name     string
	Host     string
	User     string
	Password string
	Port     string
	SSLMode  string
}

func NewSQL(cfg Config) (*sql.DB, error) {
	switch cfg.Driver {
	case DriverPostgres:
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)
		return sql.Open(cfg.Driver, connStr)
	case DriverSqlite3:
		name := cfg.Name
		if len(name) == 0 {
			name = "app_db"
		}
		return sql.Open(cfg.Driver, name)
	default:
		return nil, fmt.Errorf("invalid database driver (%s): supported drivers are postgres, sqlite3", cfg.Driver)
	}
}

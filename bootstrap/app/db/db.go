package db

import (
	"fmt"
	"log"
	"os"

	"github.com/tabajara-coder/backend-base-structure/db"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Get() *gorm.DB {
	return dbInstance
}

func init() {
	// Create a default *sql.DB exposed by the superkit/db package
	// based on the given configuration.
	config := db.Config{
		Driver:   os.Getenv("DB_DRIVER"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		Host:     os.Getenv("DB_HOST"),
	}
	dbinst, err := db.NewSQL(config)
	if err != nil {
		log.Fatal(err)
	}

	switch config.Driver {
	case db.DriverSqlite3:
		dbInstance, err = gorm.Open(sqlite.New(sqlite.Config{
			Conn: dbinst,
		}))
	case db.DriverPostgres:
		dbInstance, err = gorm.Open(postgres.New(postgres.Config{
			Conn: dbinst,
		}))
		fmt.Printf("Connected to %s\n", config.Driver)
	default:
		log.Fatal("invalid driver:", config.Driver)
	}
	if err != nil {
		log.Fatal(err)
	}
}

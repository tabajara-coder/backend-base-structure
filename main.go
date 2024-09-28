package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello, Worlddddd!")
	_, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=dev password=dev dbname=base-db port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Conectou com sucesso!")
}

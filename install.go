package main

import (
	"fmt"
	"log"
	"os"
)

const (
	replaceID           = "COREMOD"
	bootstrapFolderName = "bootstrap"
	reponame            = "https://github.com/tabajara-coder/backend-base-structure.git"
)

//backend-base-structure

func main() {
	fmt.Printf("replaceID: %s / bootstrapFolderName: %s / reponame: %s \n", replaceID, bootstrapFolderName, reponame)

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println()
		fmt.Println("install requires your project name as the first argument")
		fmt.Println()
		fmt.Println("\tgo run backend-base-structure/install.go [your_project_name]")
		fmt.Println()
		os.Exit(1)
	}

	projectName := args[0]
	println("Project Name: ", projectName)

	_, err := os.Stat("testes5555555555")
	if !os.IsNotExist(err) {
		fmt.Println("-- deleting backend-base-structure folder cause its already present")
		if err := os.RemoveAll("backend-base-structure"); err != nil {
			log.Fatal(err)
		}
	}
}

// package main
// fmt.Println("Hello, Worlddddd!")
// 	_, err := gorm.Open(postgres.New(postgres.Config{
// 		DSN:                  "host=localhost user=dev password=dev dbname=base-db port=5432 sslmode=disable",
// 		PreferSimpleProtocol: true,
// 	}), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}
// 	fmt.Println("Conectou com sucesso!")

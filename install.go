package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const (
	replaceID           = "COREMOD"
	bootstrapFolderName = "bootstrap"
	reponame            = "https://github.com/tabajara-coder/backend-base-structure.git"
)

func main() {
	projectName := getProjectNameFromCLI()

	verifyAlreadyHasDefaultProject()

	cloneProjectRepo()

	renameBaseFolderToProjectName(projectName)

	verifyProjectFilesHasReplaceID(projectName, replaceIDInProjectFiles)

	// _, err := os.Stat(projectName)
	// if !os.IsNotExist(err) {
	// 	fmt.Println("-- deleting projectName folder cause its already present")
	// 	if err := os.RemoveAll(projectName); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
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

func verifyProjectFilesHasReplaceID(projectName string, replaceFunction func(string, string, string) error) {
	err := filepath.Walk(path.Join(projectName), func(fullPath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		b, err := os.ReadFile(fullPath)
		if err != nil {
			return err
		}

		contentStr := string(b)
		if strings.Contains(contentStr, replaceID) {
			if err := replaceFunction(projectName, contentStr, fullPath); err != nil {
				return err
			}

		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func replaceIDInProjectFiles(projectName string, contentStr string, fullPath string) error {
	replacedContent := strings.ReplaceAll(contentStr, replaceID, projectName)
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(replacedContent)
	if err != nil {
		return err
	}
	return nil
}

func renameBaseFolderToProjectName(projectName string) {
	fmt.Println("-- renaming bootstrap ->", projectName)
	if err := os.Rename(path.Join("backend-base-structure", bootstrapFolderName), projectName); err != nil {
		log.Fatal(err)
	}
}

func cloneProjectRepo() {
	log.Println("-- cloning", reponame)
	clone := exec.Command("git", "clone", reponame)
	if err := clone.Run(); err != nil {
		log.Fatal(err)
	}
}

func getProjectNameFromCLI() string {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println()
		fmt.Println("install requires your project name as the first argument")
		fmt.Println()
		fmt.Println("\tgo run backend-base-structure/install.go [your_project_name]")
		fmt.Println()
		os.Exit(1)
	}
	println("\n ----Project Name: ", args[0])
	return args[0]
}

func verifyAlreadyHasDefaultProject() {
	_, err := os.Stat("backend-base-structure")
	if !os.IsNotExist(err) {
		fmt.Println("-- deleting backend-base-structure folder cause its already present")
		if err := os.RemoveAll("backend-base-structure"); err != nil {
			log.Fatal(err)
		}
	}
}

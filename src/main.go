package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

//creategoapp a cli to create go apps
//will include bin src directories
//and a main.go file in src directory

func main() {
	arg := os.Args[1]

	if len(arg) > 2 && len(arg) < 2 {

		log.Fatal("Too many arguments")
	}

	_, err := createGoAppDirectories(arg)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("cd %s\n", arg)
}

func createGoAppDirectories(arg string) (bool, error) {
	//Check if directory exist
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	appDir := cwd + "/" + arg

	if checkIfFolderExists(appDir) == false {
		srcDir := appDir + "/" + "src"
		binDir := appDir + "/" + "bin"
		createDirectories(appDir)
		createDirectories(srcDir)
		createDirectories(binDir)
		createMainFile(srcDir)
		return true, nil
	}
	return false, fmt.Errorf("%s already exists", appDir)

}

func createMainFile(basePath string) error {
	mainFile := basePath + "/" + "main.go"
	file, err := os.Create(mainFile)

	if err != nil {
		return errors.New("Unable to create main.go")
	}
	defer file.Close()
	return nil
}

func createDirectories(basePath string) error {

	err := os.Mkdir(basePath, 0755)

	if err != nil {

		return fmt.Errorf("%s", err)
	}
	return nil
}

//Check if Directory exists
func checkIfFolderExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

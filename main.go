package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// TODO: check if git binary present
	// TODO: load args from CLI
	// TODO: read config
	newAppName := "migrations"

	createAppDir(newAppName)
	addAppToGitIgnore(newAppName)
	//fmt.Println("Add container to docker-compose.apps.yml")
	goToAppDir(newAppName)
	createReadme(newAppName)
	createDockerfile(newAppName)
	initializeGitRepo(newAppName)
}

func createAppDir(newAppName string) {
	err := os.Mkdir(newAppName, 0755)
	if err != nil {
		// Check if directory exists
		fmt.Println("Can't create dir")
	}
}

func addAppToGitIgnore(newAppName string) {
	//TODO: check if this line exists in file
	wd, _ := os.Getwd()
	f, err := os.OpenFile(wd + "/.gitignore", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(newAppName + "\n")
	if err != nil {
		fmt.Println(err)
	}
	f.Close()
}

func goToAppDir(newAppName string) {
	err := os.Chdir(newAppName)
	if err != nil {
		fmt.Println(err)
	}
}

func createReadme(newAppName string) {
	readmeFile, err := os.Create("README.md")
	if err != nil {
		fmt.Println(err)
	}
	readmeFile.Close()
}

func initializeGitRepo(newAppName string) {
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	cmd = exec.Command("git", "add", ".")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	cmd = exec.Command("git", "commit", "-am", "\"Initial commit\"")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func createDockerfile(newAppName string) {
	dockerfile, err := os.Create("Dockerfile")
	if err != nil {
		fmt.Println(err)
	}
	dockerfile.Close()
}

func createBuildScript(newAppName string) {
	buildScript, err := os.Create("build.sh")
	if err != nil {
		fmt.Println(err)
	}
	buildScript.Chmod(0755)
	buildScript.Close()
}

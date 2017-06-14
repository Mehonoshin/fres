package structure

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateAppDir(newAppName string) {
	err := os.Mkdir(newAppName, 0755)
	if err != nil {
		// Check if directory exists
		fmt.Println("Can't create dir")
	}
}

func AddAppToGitIgnore(newAppName string) {
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

func CommitToMasterRepo(newAppName string) {
	// TODO: use git lib for dealing with repos
	cmd := exec.Command("git", "add", ".gitignore")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	cmd = exec.Command("git", "commit", "-am", "\"Add " + newAppName + " app\"")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func GoToAppDir(newAppName string) {
	err := os.Chdir(newAppName)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateReadme(newAppName string) {
	readmeFile, err := os.Create("README.md")
	if err != nil {
		fmt.Println(err)
	}
	readmeFile.Close()
}

func InitializeGitRepo(newAppName string) {
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

func CreateDockerfile(newAppName string) {
	dockerfile, err := os.Create("Dockerfile")
	if err != nil {
		fmt.Println(err)
	}
	dockerfile.Close()
}

func CreateBuildScript(newAppName string) {
	buildScript, err := os.Create("build.sh")
	if err != nil {
		fmt.Println(err)
	}
	buildScript.Chmod(0755)
	buildScript.Close()
}

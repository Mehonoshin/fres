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
	runShellCmd("git", "add", ".gitignore")
	runShellCmd("git", "commit", "-am", "\"Add " + newAppName + " app\"")
	//git remote add origin ssh://git@bitbucket.org/mexx/test1.git
	//git push -u origin master
}

func GoToAppDir(newAppName string) {
	err := os.Chdir(newAppName)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateReadme(newAppName string) {
	createFile("README.md")
}

func InitializeGitRepo(newAppName string) {
	runShellCmd("git", "init")
	runShellCmd("git", "add", ".")
	runShellCmd("git", "commit", "-am", "\"Initial commit\"")
}

func CreateDockerfile(newAppName string) {
	createFile("Dockerfile")
}

func CreateBuildScript(newAppName string) {
	createFile("build.sh")
}

func runShellCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func createFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	file.Chmod(0755)
	file.Close()
}

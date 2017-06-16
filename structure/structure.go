package structure

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Mehonoshin/fres/git"
	"github.com/Mehonoshin/fres/shell"
)

func CreateAppDir(newAppName string) {
	err := os.Mkdir(newAppName, 0755)
	if err != nil {
		// Check if directory exists
		fmt.Println("Can't create dir")
	}
}

func AddAppToGitIgnore(newAppName string) {
	git.AddToGitIgnore(newAppName)
}

func CommitToMasterRepo(newAppName string) {
	git.CommitFile(".gitignore", "\"Add " + newAppName + " app\"")
}

func SetupGit(newAppName string) {
	shell.RunShellCmd("git", "init")
	shell.RunShellCmd("git", "add", ".")
	shell.RunShellCmd("git", "commit", "-am", "\"Initial commit\"")

	git.CreateRemote(newAppName)
	git.AddRemoteAsOrigin(newAppName)
	git.PushMaster()
}

func GoToAppDir(newAppName string) {
	err := os.Chdir(newAppName)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateReadme(newAppName string) {
	// TODO: add default content
	createFile("README.md")
}

func CreateDockerfile(newAppName string) {
	// TODO: add default content
	createFile("Dockerfile")
}

func CreateBuildScript(newAppName string) {
	// TODO: add default content
	createFile("build.sh")
}

func createFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	file.Chmod(0755)
	file.Close()
}

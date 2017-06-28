package structure

import (
	"github.com/Mehonoshin/fres/git"
	"github.com/Mehonoshin/fres/shell"
	"github.com/Mehonoshin/fres/utils"
)

func CreateAppDir(newAppName string) {
	// return error
	utils.CreateDir(newAppName)
}

func CreateProjectDir(dirPath string) {
	// return error
	utils.CreateDir(dirPath)
}

func CreateConfig(dirPath string) {
	// parse string and return only last token
	filePath := dirPath + "/.fres.yml"
	utils.CreateFile(filePath)
	utils.WriteToFile(filePath, sampleConfig())
}

func AddAppToGitIgnore(newAppName string) {
	// return error
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

func CreateReadme(newAppName string) {
	// TODO: add default content
	utils.CreateFile("README.md")
}

func CreateDockerfile(newAppName string) {
	// TODO: add default content
	utils.CreateFile("Dockerfile")
}

func CreateBuildScript(newAppName string) {
	// TODO: add default content
	utils.CreateFile("build.sh")
}

func sampleConfig() string {
	return "project_name: myproject\n" +
	"scm: bitbucket\n" +
	"bitbucket:\n" +
	"  user: sample_user\n" +
	"  password: sample_pass\n"
}

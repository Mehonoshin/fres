package structure

import (
	"github.com/Mehonoshin/fres/git"
	"github.com/Mehonoshin/fres/utils"
	"github.com/Mehonoshin/fres/config"
)

func CreateAppDir(newAppName string) {
	// return error
	utils.CreateDir(newAppName)
}

func CreateProjectDir(dirPath string) {
	// return error
	utils.CreateDir(dirPath)
}

func CreateConfig(dirPath string) (string, error) {
	// parse string and return only last token
	filePath := dirPath + "/.fres.yml"
	utils.CreateFile(filePath)
	err := utils.WriteToFile(filePath, config.SampleConfig())
	return filePath, err
}

func AddToGitIgnore(gitignorePath, pathToAdd string) {
	// return error
	git.AddToGitIgnore(gitignorePath, pathToAdd)
}

func CommitToMasterRepo(newAppName string) {
	git.CommitFile(".gitignore", "\"Add " + newAppName + " app\"")
}

func SetupGit(newAppName string) {
	git.CreateLocalGitRepo()
	git.CreateRemote(newAppName)
	git.AddRemoteAsOrigin(newAppName)
	git.PushMaster()
}

func CreateGitRepo(path string) {
	utils.GoToDir(path)
	git.CreateLocalGitRepo()
}

func CreateReadme(newAppName string) {
	// TODO: add default content
	utils.CreateFile("README.md")
}

func CreateBuildScript(newAppName string) {
	// TODO: add default content
	utils.CreateFile("build.sh")
}

//func SampleDockerfile() string {
//}

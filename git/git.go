package git

import (
	"fmt"
	"os"

	"github.com/Mehonoshin/fres/shell"
	"github.com/Mehonoshin/fres/bitbucket"
	"github.com/Mehonoshin/fres/config"
	"github.com/Mehonoshin/fres/utils"
)


func CreateLocalGitRepo() {
	shell.RunCmd("git", "init")
	shell.RunCmd("git", "add", ".")
	shell.RunCmd("git", "commit", "-am", "\"Initial commit\"")
}

func AddToGitIgnore(gitignorePath, filename string) {
	//TODO: check if this line exists in file
	wd, _ := os.Getwd()
	f, err := os.OpenFile(wd + "/" + gitignorePath + "/.gitignore", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(filename + "\n")
	if err != nil {
		fmt.Println(err)
	}
	f.Close()
}

func CommitFile(filename, commitMessage string) {
	// TODO: use git lib for dealing with repos
	shell.RunCmd("git", "add", filename)
	shell.RunCmd("git", "commit", "-m", commitMessage)
}

func CreateRemote(appName string) {
	appName = config.Conf.ProjectName + "-" + appName

	err := bitbucket.CreateRemoteRepo(appName,
		config.Conf.Bitbucket.User,
		config.Conf.Bitbucket.AppPassword,
		config.Conf.Bitbucket.ProjectMode)

	if err != nil {
		utils.Error("Can't create '" + appName + "' remote")
	}
}

func AddRemoteAsOrigin(appName string) {
	remoteName := config.Conf.ProjectName + "-" + appName
	shell.RunCmd("git", "remote", "add", "origin", "ssh://git@bitbucket.org/" + config.Conf.Bitbucket.User + "/" + remoteName + ".git")
}

func PushMaster() {
	pushBranch("master")
}

func pushBranch(branchName string) {
	shell.RunCmd("git", "push", "-u", "origin", branchName)
}

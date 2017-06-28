package git

import (
	"fmt"
	"os"

	"github.com/Mehonoshin/fres/shell"
	"github.com/Mehonoshin/fres/bitbucket"
	"github.com/Mehonoshin/fres/config"
)

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
	shell.RunShellCmd("git", "add", filename)
	shell.RunShellCmd("git", "commit", "-m", commitMessage)
}

func CreateRemote(appName string) {
	bitbucket.CreateRemoteRepo(appName, config.Conf.Bitbucket.User, config.Conf.Bitbucket.AppPassword)
}

func AddRemoteAsOrigin(remoteName string) {
	// TODO: get butbucket username from config
	//shell.RunShellCmd("git", "remote", "add", "origin", "ssh://git@bitbucket.org/mexx/" + remoteName + ".git")
}

func PushMaster() {
	pushBranch("master")
}

func pushBranch(branchName string) {
	shell.RunShellCmd("git", "push", "-u", "origin", branchName)
}

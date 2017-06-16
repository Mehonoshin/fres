package bitbucket

import (
	"github.com/Mehonoshin/fres/shell"
)

func CreateRemoteRepo(repoName, username, appPassword string) {
	// TODO: remove error if repo not created
	// TODO: check if repo exists
	cmd := "curl --user" + username + ":" + appPassword + "-X POST -H \"Content-Type: application/json\" -d '{\"scm\": \"git\"}' https://api.bitbucket.org/2.0/repositories/" + username + "/" + repoName
	shell.RunShellCmd(cmd)
}

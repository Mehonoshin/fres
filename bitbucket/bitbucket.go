package bitbucket

import (
	"github.com/Mehonoshin/fres/utils"
)

func CreateRemoteRepo(repoName, username, appPassword string) error {
	// TODO: remove error if repo not created
	// TODO: check if repo exists
	endpoint := "https://api.bitbucket.org/2.0/repositories/" + username + "/" + repoName
	payload := "{\"scm\": \"git\"}"

	err := utils.HttpPostJson(endpoint, username, appPassword, payload)

	return err
}

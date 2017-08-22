package bitbucket

import (
	"github.com/Mehonoshin/fres/utils"
)

func CreateRemoteRepo(repoName, username, appPassword string, project_mode bool) error {
	// TODO: remove error if repo not created
	// TODO: check if repo exists
	endpoint := "https://api.bitbucket.org/2.0/repositories/" + username + "/" + repoName
	payload := "{\"scm\": \"git\"}"

	// TODO: check response and handle errors
	err, _ := utils.HttpPostJson(endpoint, username, appPassword, payload)

	if err != nil {
		utils.Error(err)
	}

	return err
}

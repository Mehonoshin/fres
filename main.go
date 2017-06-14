package main

import (
	_ "fmt"
	"github.com/Mehonoshin/fres/structure"
)

// TODO: implement the following features
// fres init [project name]
// fres create [app name] --lang ruby
// fres remove [app name]
// fres deploy [app name]
// fres deploy

func main() {
	// TODO: check if git binary present
	// TODO: load args from CLI
	// TODO: read config
	// TODO: display https://developer.atlassian.com/bitbucket/api/2/reference/meta/authentication app passwords text if no pass provided
	newAppName := "migrations"

	structure.CreateAppDir(newAppName)
	structure.AddAppToGitIgnore(newAppName)
	structure.CommitToMasterRepo(newAppName)
	//fmt.Println("Add container to docker-compose.apps.yml")
	structure.GoToAppDir(newAppName)
	structure.CreateReadme(newAppName)
	structure.CreateDockerfile(newAppName)
	structure.CreateBuildScript(newAppName)
	structure.InitializeGitRepo(newAppName)
	// TODO: create remote repo and push
}



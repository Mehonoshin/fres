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
	newAppName := "migrations"

	structure.CreateAppDir(newAppName)
	structure.AddAppToGitIgnore(newAppName)
	structure.CommitToMasterRepo(newAppName)
	//fmt.Println("Add container to docker-compose.apps.yml")
	structure.GoToAppDir(newAppName)
	structure.CreateReadme(newAppName)
	structure.CreateDockerfile(newAppName)
	structure.InitializeGitRepo(newAppName)
	// TODO: create remote repo and push
}



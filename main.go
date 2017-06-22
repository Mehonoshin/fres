package main

import (
	"fmt"

	"github.com/Mehonoshin/fres/structure"
	"github.com/Mehonoshin/fres/utils"
	"github.com/Mehonoshin/fres/config"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose    = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	configPath = kingpin.Flag("config", "Path to config").Default(".fres.yml").String()

	cmd				 = kingpin.Arg("cmd", "Command to perform.").Required().String()
	name			 = kingpin.Arg("name", "Name").Default(".").String()
)

// TODO: implement the following features
// fres init [project name]
// fres create [app name] --lang ruby
// fres remove [app name]
// fres deploy [app name]
// fres deploy

func main() {
	kingpin.Parse()
	config.Load(*configPath)

	switch *cmd {
		case "init":
			utils.Message("init")
		case "create":
			utils.Message("create")
			create(*name)
		case "remove":
			utils.Message("remove")
		case "deploy":
			utils.Message("deploy")
		default:
			utils.Error("Unknown command")
	}
}

func initProject(initPath string) {
	//init.CreateProjectDir(initPath)
	//init.AddConfigToGitIgnore()
	//init.CreateConfig()
}

func create(newAppName string) {
	// TODO: check if git binary present
	// TODO: load args from CLI
	// TODO: read config
	// TODO: display https://developer.atlassian.com/bitbucket/api/2/reference/meta/authentication app passwords text if no pass provided

	// Master dir
	fmt.Println(config.Conf)
	structure.CreateAppDir(newAppName)
	structure.AddAppToGitIgnore(newAppName)
	structure.CommitToMasterRepo(newAppName)
	fmt.Println("Add container to docker-compose.apps.yml")

	// New Project dir
	structure.GoToAppDir(newAppName)
	structure.CreateReadme(newAppName)
	structure.CreateDockerfile(newAppName)
	structure.CreateBuildScript(newAppName)
	structure.SetupGit(newAppName)
}

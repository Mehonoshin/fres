package main

import (
	"github.com/Mehonoshin/fres/structure"
	"github.com/Mehonoshin/fres/utils"
	"github.com/Mehonoshin/fres/config"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose    = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	configPath = kingpin.Flag("config", "Path to config").Default(".fres.yml").String()
	lang       = kingpin.Flag("lang", "Service language").Short('l').Default("ruby").String()

	cmd				 = kingpin.Arg("cmd", "Command to perform(init, create, remove, deploy)").Required().String()
	name			 = kingpin.Arg("name", "Name").Default(".").String()
)

func main() {
	kingpin.Parse()

	// TODO: check if git binary present
	// TODO: display https://developer.atlassian.com/bitbucket/api/2/reference/meta/authentication app passwords text if no pass provided

	loadConfig(*cmd)

	switch *cmd {
		case "init":
			utils.Message("Initialize new project '" + *name + "'")
			initProject(*name)
		case "create":
			utils.Message("Create new service '" + *name + "'")
			create(*name)
		case "remove":
			utils.Message("remove")
		case "deploy":
			utils.Message("deploy")
		default:
			utils.Error("Unknown command")
	}
}

func loadConfig(command string) {
	if (command != "init") {
		config.Load(*configPath)
	} else {
		config.Conf = &config.Config{}
	}

	config.Conf.Cmd = *cmd
	config.Conf.Arg = *name
}

func initProject(initPath string) {
	structure.CreateProjectDir(initPath)

	_, err := structure.CreateConfig(initPath)
	if err != nil {
		utils.Warn(err)
	}

	utils.CreateFile(initPath + "/.gitignore")
	structure.AddToGitIgnore(initPath, ".fres.yml")
	structure.CreateGitRepo(initPath)
}

func create(newAppName string) {
	// Master dir
	structure.CreateAppDir(newAppName)
	structure.AddToGitIgnore(".", newAppName)
	structure.CommitToMasterRepo(newAppName)
	utils.Message("Add container to docker-compose.apps.yml")

	// New Project dir
	utils.GoToDir(newAppName)
	structure.CreateReadme(newAppName)
	structure.CreateDockerfile(newAppName)
	structure.CreateBuildScript(newAppName)
	structure.SetupGit(newAppName)
	utils.Success(newAppName + " successfully created")
}

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
	lang       = kingpin.Flag("lang", "Service language").Short('l').Default("ruby").String()

	cmd				 = kingpin.Arg("cmd", "Command to perform.").Required().String()
	name			 = kingpin.Arg("name", "Name").Default(".").String()
)

func main() {
	kingpin.Parse()

	if (*cmd != "init") {
		loadConfig()
	}

	switch *cmd {
		case "init":
			utils.Message("init")
			initProject(*name)
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

func loadConfig() {
	config.Load(*configPath)
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
	// TODO: check if git binary present
	// TODO: load args from CLI
	// TODO: read config
	// TODO: display https://developer.atlassian.com/bitbucket/api/2/reference/meta/authentication app passwords text if no pass provided

	// Master dir
	fmt.Println(config.Conf)
	structure.CreateAppDir(newAppName)
	structure.AddToGitIgnore(".", newAppName)
	structure.CommitToMasterRepo(newAppName)
	fmt.Println("Add container to docker-compose.apps.yml")

	// New Project dir
	utils.GoToDir(newAppName)
	structure.CreateReadme(newAppName)
	structure.CreateDockerfile(newAppName)
	structure.CreateBuildScript(newAppName)
	structure.SetupGit(newAppName)
}

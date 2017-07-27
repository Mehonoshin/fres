package main

import (
	"github.com/Mehonoshin/fres/structure"
	"github.com/Mehonoshin/fres/utils"
	"github.com/Mehonoshin/fres/config"
	"github.com/Mehonoshin/fres/shell"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose     = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	configPath  = kingpin.Flag("config", "Path to config").Default(".fres.yml").String()
	lang        = kingpin.Flag("lang", "Service language").Default("ruby").String()

	vcs         = kingpin.Flag("vcs", "Version control system").Default("git").String()
	vcsProvider = kingpin.Flag("vcs-provider", "Provider of version control system(github, bitbucket)").Default("bitbucket").String()

	cmd				 = kingpin.Arg("cmd", "Command to perform(init, create, remove, deploy)").Required().String()
	name			 = kingpin.Arg("name", "Name").Default(".").String()
)

func main() {
	kingpin.Parse()

	// TODO: check if required packages are installed(git, mercurial)

	loadConfig(*cmd)
	runCommand()
}

func loadConfig(command string) {
	if (command != "init") {
		config.Load(*configPath)
	} else {
		config.Conf = &config.Config{}
	}

	config.Conf.Runtime = &config.RuntimeConfig{
		Cmd: *cmd,
		Arg: *name,
		Vcs: *vcs,
		VcsProvider: *vcsProvider,
		Lang: *lang,
	}
}

func runCommand() {
	switch *cmd {
		case "init":
			utils.Message("Initialize new project '" + *name + "'")
			initProject(*name)
		case "create":
			// TODO: check if config exists before creating new service
			// TODO: Create gemfile for ruby services
			// TODO: populate build.sh script
			utils.Message("Create new service '" + *name + "'")
			create(*name)
		case "remove":
			utils.Message("remove")
		case "deploy":
			utils.Message("Deploy whole project")
			deploy()
		default:
			utils.Error("Unknown command")
	}
}

func initProject(initPath string) {
	// TODO: replace initPath with project root
	// if bitbucket chosen
	// TODO: display https://developer.atlassian.com/bitbucket/api/2/reference/meta/authentication app passwords text if no pass provided
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
	CreateDockerfile(newAppName)
	structure.CreateBuildScript(newAppName)
	structure.SetupGit(newAppName)
	utils.Success(newAppName + " successfully created")
}

func deploy() {
	deployConf := config.Conf.Deploy
	utils.GoToDir(deployConf.Root)
	shell.RunCmd(deployConf.Cmd)
}

package config

import (
	"io/ioutil"
	"log"

	_ "github.com/Mehonoshin/fres/utils"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Runtime *RuntimeConfig

	ProjectName string   `yaml:"project_name"`
	Vcs         string   `yaml:"vcs"`
	VcsProvider string   `yaml:"vcs_provider"`

	Bitbucket   struct {
		User        string `yaml:"user"`
		AppPassword string `yaml:"app_password"`
		ProjectMode bool   `yaml:"project_mode"`
	}

	Languages []struct {
		Name      string `yaml:"name"`
		Version   string `yaml:"version"`
		BaseImage string `yaml:"base_image"`
	}

	Deploy struct {
		Root string `yaml:"root"`
		Cmd  string `yaml:"cmd"`
	}
}

type RuntimeConfig struct {
	Cmd string
	Arg string

	Vcs         string
	VcsProvider string

	Lang string
}

var (
	Conf *Config
)

func NewConfig() *Config {
	c := &Config{}
	c.Runtime = &RuntimeConfig{}
	return c
}

func Load(path string) *Config {
	c := NewConfig()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Lookup config")
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		// TODO: replace with utils.Error
		log.Printf("Can't parse config:", err)
		c = NewConfig()
	}

	Conf = c

	return c
}

//func lookupProjectRoot()

func SampleConfig() string {
	return "project_name: " + Conf.Runtime.Arg + "\n" +
	"vcs: " + Conf.Runtime.Vcs + "\n" +
	"vcs_provider: " + Conf.Runtime.VcsProvider + "\n" +
	"\n" +
	"bitbucket:\n" +
	"  user: sample_user\n" +
	"  password: sample_pass\n"
}

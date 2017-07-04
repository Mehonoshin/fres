package config

import (
	"io/ioutil"
	"log"

	_ "github.com/Mehonoshin/fres/utils"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RuntimeConfig

	ProjectName string   `yaml:"project_name"`
	Scm         string   `yaml:"scm"`

	Bitbucket   struct {
		User        string `yaml:"user"`
		AppPassword string `yaml:"app_password"`
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
}

var (
	Conf *Config
)

func Load(path string) *Config {
	c := &Config{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Can't read config:", err)
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		log.Printf("Can't parse config:", err)
		c = &Config{}
	}

	Conf = c

	return c
}

func SampleConfig() string {
	return "project_name: " + Conf.Arg + "\n" +
	"scm: bitbucket\n" +
	"bitbucket:\n" +
	"  user: sample_user\n" +
	"  password: sample_pass\n"
}

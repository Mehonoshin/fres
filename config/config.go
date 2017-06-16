package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ProjectName string   `yaml:"project_name"`
	Scm         string   `yaml:"scm"`
	Bitbucket   struct {
		User        string `yaml:"user"`
		AppPassword string `yaml:"app_password"`
	}
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
	}

	Conf = c

	return c
}

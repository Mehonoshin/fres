package main

import (
	"github.com/Mehonoshin/fres/utils"
	"github.com/Mehonoshin/fres/config"

	_ "github.com/valyala/fasttemplate"
)

func CreateDockerfile(newAppName string) {
	// TODO: add default content
	filePath := "Dockerfile"
	utils.CreateFile(filePath)
	utils.WriteToFile(filePath, dockerfileContent(config.Conf.Runtime.Lang))
}

func dockerfileContent(lang string) string {
	return `FROM {{base_image}}\n
MAINTAINER {{maintainer}} <{{maintainer_email}}>\n` +
	languageSpecificContent(lang)
}

func languageSpecificContent(lang string) string {
	switch lang {
	case "ruby":
		return rubyDockerfile()
	default:
		return "\n"
	}
}

// TODO: add Elixir support
// TODO: unescape line breaks
// TODO: replace variables in templates

func rubyDockerfile() string {
	return `COPY . /app/ \n
WORKDIR /app \n
RUN bundle install \n
CMD foreman start`
}

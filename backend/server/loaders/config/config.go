package config

import (
	"os"

	"share/utils/text"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"server/utils/logger"
)

var C = new(Model)

func init() {
	// Load configurations to struct
	yml, err := os.ReadFile("config.yaml")
	if err != nil {
		logger.Log(logrus.Fatal, "UNABLE TO READ YAML CONFIGURATION FILE")
	}
	if err := yaml.Unmarshal(yml, C); err != nil {
		logger.Log(logrus.Fatal, "UNABLE TO PARSE YAML CONFIGURATION FILE")
	}

	// Validate configurations
	if err := text.Validator.Struct(C); err != nil {
		logger.Log(logrus.Fatal, "INVALID CONFIGURATION: "+err.Error())
	}

	// Apply log level configuration
	logrus.SetLevel(logrus.Level(C.LogLevel))
	spew.Config = spew.ConfigState{Indent: "  "}
}

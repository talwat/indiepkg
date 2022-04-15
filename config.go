package main

import (
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var mainPath string = home + ".indiepkg/"
var srcPath string = mainPath + "data/package_src/"
var infoPath string = mainPath + "data/installed_packages/"
var configPath string = mainPath + "config/"
var binPath string

type Config struct {
	Paths struct {
		Prefix string
	}
}

var config Config

func loadConfig() {
	log(1, "Reading config file...")
	raw := readFile(configPath+"config.toml", "An error occurred while reading config file")
	log(1, "Loading config file...")
	err := toml.Unmarshal([]byte(raw), &config)
	errorLog(err, 4, "An error occurred while loading config file")

	config.Paths.Prefix = home + config.Paths.Prefix
	if !strings.HasSuffix(config.Paths.Prefix, "/") {
		config.Paths.Prefix += "/"
	}
	newDir(config.Paths.Prefix, "An error occurred while creating prefix directory")

	binPath = config.Paths.Prefix + "bin/"
}

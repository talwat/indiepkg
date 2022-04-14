package main

import (
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var mainPath string = home + ".indiepkg/"
var binPath string = home + ".local/bin/"
var prefixPath string = home + ".local/"
var srcPath string = mainPath + "data/package_src/"
var infoPath string = mainPath + "data/installed_packages/"
var configPath string = mainPath + "config/"

type Config struct {
	Paths struct {
		Bin    string
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

	config.Paths.Bin = home + config.Paths.Bin
	if !strings.HasSuffix(config.Paths.Bin, "/") {
		config.Paths.Bin += "/"
	}
	newDir(config.Paths.Bin, "An error occurred while creating binary directory")

	config.Paths.Prefix = home + config.Paths.Prefix
	if !strings.HasSuffix(config.Paths.Prefix, "/") {
		config.Paths.Prefix += "/"
	}
	newDir(config.Paths.Bin, "An error occurred while creating prefix directory")
}

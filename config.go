package main

import (
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var (
	mainPath        string = home + ".indiepkg/"
	srcPath         string = mainPath + "data/package_src/"
	tmpSrcPath      string = mainPath + "tmp/package_src/"
	infoPath        string = mainPath + "data/installed_packages/"
	configPath      string = mainPath + "config/"
	tmpPath         string = mainPath + "tmp/"
	indiePkgSrcPath string = mainPath + "src/"
	indiePkgBin     string = home + ".local/bin/indiepkg"
)

type Paths struct {
	Prefix string
}

type Updating struct {
	Branch     string
	AutoUpdate bool
}

type Github struct {
	Username string
	Token    string
}

type Progressbar struct {
	Saucer        string
	SaucerHead    string
	AltSaucerHead string
	SaucerPadding string
	BarStart      string
	BarEnd        string
}

type Config struct {
	Paths Paths

	Updating Updating

	Progressbar Progressbar

	Github Github
}

var config Config = Config{
	Paths{
		home + ".local/",
	},

	Updating{
		"testing",
		true,
	},

	Progressbar{
		"[cyan]=[reset]",
		"[cyan]>[reset]",
		"[cyan]>[reset]",
		" ",
		"(",
		")",
	},

	Github{
		"username",
		"token",
	},
}

func loadConfig() {
	log(1, "Reading config file...")

	raw := readFile(configPath+"config.toml", "An error occurred while reading config file")

	log(1, "Loading config file...")

	err := toml.Unmarshal([]byte(raw), &config)

	errorLog(err, "An error occurred while loading config file")

	if !strings.HasPrefix(config.Paths.Prefix, home) { // If not in the home directory, prepend it
		config.Paths.Prefix = home + config.Paths.Prefix
	}

	if !strings.HasSuffix(config.Paths.Prefix, "/") { // If doesn't end with a /, add one
		config.Paths.Prefix += "/"
	}
}

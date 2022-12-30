package main

import (
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var (
	mainPath        string = chooseDataDir()
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

type Progress struct {
	CompileProgressIndicator string `toml:"compile_progress_indicator"`
}

type Config struct {
	Paths Paths

	Updating Updating

	Progressbar Progressbar

	Github Github

	Progress Progress
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

	Progress{
		"spinner",
	},
}

// Loads & reads configuration file.
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

// chooseDataDir returns the best directory to store data based on $INDIEPKG_DATADIR and $XDG_DATA_HOME,
// using ~/.local/share/indiepkg as the default if neither is set.
func chooseDataDir() string {
	home, _ := os.UserHomeDir()
	remEnv := os.Getenv("INDIEPKG_DATADIR")
	if remEnv != "" {
		return remEnv
	}
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome != "" {
		return dataHome + "/indiepkg/"
	}
	return home + "/.local/share/indiepkg/"
}

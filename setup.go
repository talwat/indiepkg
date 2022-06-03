package main

import "os"

// Setup directories & source.
func setup() {
	log(1, "IndiePKG initial setup.")
	chapLog("=>", "", "Initializing directories & source")
	initDirs(false)

	chapLog("=>", "", "Updating & recompiling IndiePKG")
	pullSrcRepo(false)
	compSrc()

	log(1, "Initialized and updated IndiePKG.")
}

// Set environment variables in shell rc files.
func envAdd() {
	confirm("y", "Are you sure you would like to add %s to several environment variables? This will fix a lot of issues with packages not being found (y/n)", bolden(home+".local"))

	chapLog("=>", "", "Initializing")
	newDir(configPath, "An error occurred while creating config directory")
	safeNewFile(configPath+"config.toml", "config file", false, defaultConf)
	loadConfig()

	appendVarRc := func(varName string, path string) string {
		textToAppend := "export " + varName + "=\"" + path + ":$" + varName + "\"" + "\n"

		return textToAppend
	}

	appendRc := func(name string, text string) {
		appendToFile(home+"."+name, text)
	}

	fullAppendRc := func(name string) {
		log(1, "Checking if %s exists...", bolden(name))

		if !pathExists(home+"."+name, name) {
			return
		}

		log(1, "%s found, adding to it...", bolden(name))

		appendRc(
			name,
			"\n\n# Start of IndiePKG additions\n"+
				appendVarRc("PATH", config.Paths.Prefix+"bin")+
				appendVarRc("CPATH", config.Paths.Prefix+"include")+
				appendVarRc("LD_LIBRARY_PATH", config.Paths.Prefix+"lib")+
				appendVarRc("PKG_CONFIG_PATH", config.Paths.Prefix+"pkgconfig")+
				"# End of IndiePKG additions",
		)
		log(0, "Appended to %s.", bolden(name))
	}

	chapLog("=>", "", "Appending")
	fullAppendRc("bashrc")
	fullAppendRc("zshrc")
	fullAppendRc("kshrc")
	fullAppendRc("profile")
	fullAppendRc("zprofile")

	if checkIfCommandExists("fish") {
		log(3, "Fish found, but fish is not supported officially. You will have to add the environment variables manually.")
	}

	chapLog("=>", textCol.Green, "Success")
	log(0, "Done! Restart your shell by running %s", bolden("exec "+os.Getenv("SHELL")))
}

package main

import "os"

func setup() {
	log(1, "IndiePKG initial setup.")
	chapLog("=>", "", "Initializing directories & source")
	initDirs(false)

	chapLog("=>", "", "Updating & recompiling IndiePKG")
	pullSrcRepo(false)
	compSrc()

	log(1, "Initialized and updated IndiePKG.")
}

func envAdd() {
	confirm("y", "Are you sure you would like to add %s to several environment variables? This will fix a lot of issues with packages not being found (y/n)", bolden(home+".local"))

	appendVarRc := func(varName string, path string) string {
		textToAppend := "export " + varName + "=\"$HOME/" + path + ":$" + varName + "\"" + "\n"

		return textToAppend
	}

	appendRc := func(name string, text string) {
		appendToFile(home+"."+name, text)
	}

	fullAppendRc := func(name string) {
		if !pathExists(home+"."+name, name) {
			return
		}

		log(1, "%s found, adding to it...", bolden(name))

		appendRc(
			name,
			"\n\n# Start of IndiePKG additions\n"+
				appendVarRc("PATH", ".local/bin")+
				appendVarRc("CPATH", ".local/include")+
				appendVarRc("LD_LIBRARY_PATH", ".local/lib")+
				appendVarRc("PKG_CONFIG_PATH", ".local/lib/pkgconfig")+
				"# End of IndiePKG additions",
		)
		log(0, "Appended to %s.", bolden(name))
	}

	fullAppendRc("bashrc")
	fullAppendRc("zshrc")
	fullAppendRc("kshrc")
	fullAppendRc("profile")
	fullAppendRc("zprofile")

	if checkIfCommandExists("fish") {
		log(3, "Fish found, but fish is not supported officially. You will have to add the environment variables manually.")
	}

	log(0, "Done! Restart your shell by running %s", bolden("exec "+os.Getenv("SHELL")))
}

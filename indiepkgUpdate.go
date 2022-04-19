package main

func updateIndiePKG() {
	chapLog("=>", "", "Initializing")
	initDirs(false)
	loadConfig()

	chapLog("=>", "", "Updating IndiePKG")
	chapLog("==>", "", "Pulling source code")
	if pullSrcRepo(false) {
		return
	}

	chapLog("==>", "", "Compiling IndiePKG")
	logNoNewline(1, "Running %s", bolden("make"))
	runCommandDot(indiePkgSrcDir, "make")

	chapLog("==>", "", "Moving IndiePKG binary")
	mvPath(indiePkgSrcDir+"indiepkg", home+".local/bin/indiepkg")

	chapLog("=>", "GREEN", "Success")
	log(0, "Updated IndiePKG!")
}

func autoUpdate() {
	if config.Updating.Auto_update {
		log(1, "Checking for an update...")

		if pullSrcRepo(true) {
			return
		}

		runCommand(indiePkgSrcDir, "make")
		mvPath(indiePkgSrcDir+"indiepkg", home+".local/bin/indiepkg")
		log(0, "Auto-updated IndiePKG!")
	}
}

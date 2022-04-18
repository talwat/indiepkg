package main

func updateIndiePKG() {
	chapLog("=>", "", "Initializing")
	initDirs(false)
	loadConfig()

	chapLog("=>", "", "Updating IndiePKG")
	chapLog("==>", "", "Pulling source code")
	pullSrcRepo(false) //nolint:errcheck

	chapLog("==>", "", "Compiling IndiePKG")
	logNoNewline(1, "Running %s", bolden("make"))
	runCommandRealTime(indiePkgSrcDir, "make")

	chapLog("==>", "", "Moving IndiePKG binary")
	mvPath(indiePkgSrcDir+"indiepkg", home+".local/bin/indiepkg")

	chapLog("=>", "GREEN", "Success")
	log(0, "Updated IndiePKG!")
}

func autoUpdate() {
	if config.Updating.Auto_update {
		log(1, "Checking for an update...")
		err := pullSrcRepo(true)
		if err.Error() == "already up-to-date" {
			debugLog("Auto-update returns already up to date")
			return
		}

		runCommand(indiePkgSrcDir, "make")
		mvPath(indiePkgSrcDir+"indiepkg", home+".local/bin/indiepkg")
		log(0, "Auto-updated IndiePKG!")
	}
}

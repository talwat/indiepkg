package main

func updateIndiePKG() {
	chapLog("=>", "VIOLET", "Initializing")
	initDirs(false)
	loadConfig()

	chapLog("=>", "VIOLET", "Updating IndiePKG")
	chapLog("==>", "BLUE", "Pulling source code")
	pullSrcRepo(false)

	chapLog("==>", "BLUE", "Compiling IndiePKG")
	logNoNewline(1, "Running %s", bolden("make"))
	runCommandRealTime(indiePkgSrcDir, "make")

	chapLog("==>", "BLUE", "Moving IndiePKG binary")
	mvPath(indiePkgSrcDir+"indiepkg", home+".local/bin/indiepkg")

	chapLog("=>", "GREEN", "Success")
	log(0, "Updated IndiePKG!")
}

func autoUpdate() {
	pullSrcRepo(true)
	runCommand(indiePkgSrcDir, "make")
	mvPath(indiePkgSrcDir+"indiepkg", home+".local/bin/indiepkg")
	log(0, "Auto-updated IndiePKG!")
}

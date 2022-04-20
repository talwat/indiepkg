package main

import "fmt"

func compSrc() {
	chapLog("==>", "", "Compiling IndiePKG")
	logNoNewline(1, "Running %s", bolden("make"))
	runCommandDot(indiePkgSrcDir, "make")
	fmt.Print("\n")

	chapLog("==>", "", "Moving IndiePKG binary")
	srcPath := indiePkgSrcDir + "indiepkg"
	destPath := home + ".local/bin/indiepkg"

	log(1, "Moving %s to %s...", bolden(srcPath), bolden(destPath))
	mvPath(srcPath, destPath)
}

func pullSrc() {
	chapLog("==>", "", "Pulling source code")
	if pullSrcRepo(false) {
		return
	}
}

func updateIndiePKG() {
	chapLog("=>", "", "Initializing")
	initDirs(false)
	loadConfig()

	chapLog("=>", "", "Updating IndiePKG")
	pullSrc()
	compSrc()

	chapLog("=>", "GREEN", "Success")
	log(0, "Updated IndiePKG!")
}

func autoUpdate() {
	if config.Updating.Auto_update {
		log(1, "Checking for an update...")

		if pullSrcRepo(true) {
			return
		}

		_, err := runCommand(indiePkgSrcDir, "make")
		errorLog(err, 4, "An error occurred while compiling IndiePKG because of an auto-update")
		mvPath(indiePkgSrcDir+"indiepkg", home+".local/bin/indiepkg")
		log(0, "Auto-updated IndiePKG!")
	}
}

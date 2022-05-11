package main

import (
	"os"
	"strings"
)

func compSrc() {
	chapLog("==>", "", "Compiling IndiePKG")
	logNoNewline(1, "Running %s", bolden("make"))
	runCommandDot(indiePkgSrcDir, "make")
	rawLogf("\n")

	chapLog("==>", "", "Moving IndiePKG binary")

	srcPath := indiePkgSrcDir + "indiepkg"
	destPath := Home + ".local/bin/indiepkg"

	log(1, "Moving %s to %s...", bolden(srcPath), bolden(destPath))
	mvPath(srcPath, destPath)
}

func updateIndiePKG() {
	chapLog("=>", "", "Initializing")
	initDirs(false)
	loadConfig()

	chapLog("=>", "", "Updating IndiePKG")
	chapLog("==>", "", "Pulling source code")

	if pullSrcRepo(true) {
		if !force {
			chapLog("=>", "GREEN", "Success")
			log(0, "IndiePKG already up to date.")
			os.Exit(0)
		}

		log(3, "IndiePKG already up to date, but force is on, so continuing.")
	}

	compSrc()

	chapLog("=>", "GREEN", "Success")
	log(0, "Updated IndiePKG!")
}

func autoUpdate() {
	if config.Updating.AutoUpdate {
		log(1, "Checking for an update...")

		resp, err := makeReq("http://clients3.google.com/generate_204")

		if err != nil && strings.HasSuffix(err.Error(), "no such host") {
			debugLog("Error of ping: %s", bolden(err.Error()))
			log(3, "No internet connection, skipping auto-update.")

			return
		}

		debugLog("StatusCode of ping: %s", bolden(resp.StatusCode))

		if pullSrcRepo(true) {
			return
		}

		_, err = runCommand(indiePkgSrcDir, "make")
		errorLog(err, "An error occurred while compiling IndiePKG because of an auto-update")
		mvPath(indiePkgSrcDir+"indiepkg", Home+".local/bin/indiepkg")
		log(0, "Auto-updated IndiePKG!")
	}
}

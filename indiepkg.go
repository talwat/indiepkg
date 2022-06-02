package main

import (
	"os"
	"strings"
)

// Compiles source of indiepkg
func compSrc() {
	chapLog("==>", "", "Compiling IndiePKG")
	logNoNewline(1, "Running %s", bolden("make"))
	runCommandDot(indiePkgSrcPath, false, "make")
	rawLogf("\n")

	chapLog("==>", "", "Moving IndiePKG binary")

	srcPath := indiePkgSrcPath + "indiepkg"

	log(1, "Moving %s to %s...", bolden(srcPath), bolden(indiePkgBin))
	mvPath(srcPath, indiePkgBin)
}

// Updates indiepkg
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

// Checks for an update silently
func autoUpdate() {
	if config.Updating.AutoUpdate {
		log(1, "Checking for an update...")

		resp, err := makeReq("http://clients3.google.com/generate_204") // Make request to check if offline

		if err != nil && strings.HasSuffix(err.Error(), "no such host") {
			debugLog("Error of ping: %s", bolden(err.Error()))
			log(3, "No internet connection, skipping auto-update.")

			return
		}

		debugLog("StatusCode of ping: %s", bolden(resp.StatusCode))

		if pullSrcRepo(true) {
			return
		}

		_, err = runCommand(indiePkgSrcPath, "make")
		errorLog(err, "An error occurred while compiling IndiePKG because of an auto-update")
		mvPath(indiePkgSrcPath+"indiepkg", indiePkgBin)
		log(0, "Auto-updated IndiePKG!")
	}
}

// Re-clones IndiePKG source
func reClone() {
	loadConfig()
	log(1, "Resetting IndiePKG source directory...")
	delPath(true, indiePkgSrcPath, "An error occurred while deleting the IndiePKG source directory")

	cloneSrcRepo()
	log(0, "Successfully re-cloned IndiePKG source.")
}

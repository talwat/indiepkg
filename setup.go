package main

func setup() {
	log(1, "IndiePKG initial setup.")
	chapLog("=>", "", "Initializing directories & source")
	initDirs(false)

	chapLog("=>", "", "Updating & recompiling IndiePKG")
	pullSrc()
	compSrc()

	log(1, "Initialized and updated IndiePKG.")
}

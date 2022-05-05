package main

func setup() {
	log(1, "IndiePKG initial setup.")
	chapLog("=>", "", "Initializing directories & source")
	initDirs(false)

	chapLog("=>", "", "Updating & recompiling IndiePKG")
	pullSrcRepo(false)
	compSrc()

	log(1, "Initialized and updated IndiePKG.")
}

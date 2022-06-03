package main

import (
	"strings"
)

// Gets and writes a packages info to the proper location.
func rawGetInfo(pkgName string, pkg Package) {
	log(1, "Getting info for %s", bolden(pkgName))
	log(1, "Checking for info URL...")

	if pkg.InfoURL == "" { // Check if package didn't specify info URL
		log(1, "Getting info from repos...")

		writePkg(pkgName, findPkg(pkgName))
	} else {
		log(1, "Getting info from %s...", bolden(pkg.InfoURL))

		writePkg(pkgName, getPkgFromURL(pkgName, pkg.InfoURL))
	}
}

func updatePackage(pkgNames []string) {
	fullInit()

	for _, pkgName := range pkgNames {
		chapLog("=>", "", "Updating %s", pkgName)
		pkgDisplayName := bolden(pkgName)

		if !pkgExists(pkgName) {
			log(3, "%s is not installed, so it can't be updated.", pkgDisplayName)

			continue
		}

		rawGetInfo(pkgName, readLoad(pkgName))

		chapLog("==>", textCol.Green, "Success")
		log(0, "Successfully updated package info for %s.", pkgDisplayName)
	}

	chapLog("=>", textCol.Green, "Success")
	log(0, "Successfully updated info for all selected packages.")
}

func updateAllPackages() {
	fullInit()

	chapLog("==>", "", "Getting installed packages")

	installedPkgs := make([]string, 0)

	log(1, "Getting contents of %s", bolden(infoPath))

	files := dirContents(infoPath, "An error occurred while getting list of installed packages")

	log(1, "Iterating through contents of %s", bolden(infoPath))

	for _, file := range files {
		installedPkgs = append(installedPkgs, strings.TrimSuffix(file.Name(), ".json")) // Trim .json from end of file name & append to installedPkgs
	}

	chapLog("=>", "", "Updating packages")

	for _, pkgName := range installedPkgs { // Iterate through installed packages
		chapLog("==>", "", "Updating %s", pkgName)
		rawGetInfo(pkgName, readLoad(pkgName))
		chapLog("===>", textCol.Green, "Success")
		log(0, "Successfully updated package info for %s.", bolden(pkgName))
	}

	chapLog("=>", textCol.Green, "Success")
	log(0, "Successfully updated info for all packages.")
}

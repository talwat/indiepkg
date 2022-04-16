package main

import (
	"strings"
)

func updatePackage(pkgNames []string) {
	fullInit()

	for _, pkgName := range pkgNames {
		chapLog("=>", "VIOLET", "Updating %s", pkgName)
		pkgDisplayName := bolden(pkgName)

		if !pkgExists(pkgName) {
			log(3, "%s is not installed, so it can't be updated.", pkgDisplayName)
			continue
		}

		downloadPkg(pkgName, false)

		chapLog("=>", "GREEN", "Success")
		log(0, "Successfully updated package info for %s!", pkgDisplayName)
	}
}

func updateAllPackages() {
	fullInit()

	chapLog("==>", "BLUE", "Getting installed packages")
	var installedPackages []string
	files := dirContents(infoPath, "An error occurred while getting list of installed packages")

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}

	chapLog("=>", "VIOLET", "Starting upgrades")
	for _, installedPackage := range installedPackages {
		chapLog("==>", "BLUE", "Upgrading %s", installedPackage)
		downloadPkg(installedPackage, false)
	}

	chapLog("=>", "GREEN", "Success")
	log(0, "Successfully updated info for all packages!")
}

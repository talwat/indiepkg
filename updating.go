package main

import (
	"strings"
)

func updatePackage(pkgNames []string) {
	for _, pkgName := range pkgNames {
		pkgDisplayName := bolden(pkgName)

		if !pkgExists(pkgName) {
			log(3, "%s is not installed, so it can't be updated.", pkgDisplayName)
			continue
		}

		downloadPkg(pkgName, false)

		log(0, "Successfully updated package info for %s!\n", pkgDisplayName)
	}
}

func updateAllPackages() {
	var installedPackages []string
	files := dirContents(infoPath, "An error occurred while getting list of installed packages")

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}

	log(1, "Updating all packages...")
	for _, installedPackage := range installedPackages {
		downloadPkg(installedPackage, false)
	}

	log(0, "Successfully updated info for all packages!")
}

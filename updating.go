package main

import (
	"strings"
)

func updatePackage(pkgNames []string) {
	fullInit()

	for _, pkgName := range pkgNames {
		chapLog("=>", "", "Updating %s", pkgName)
		pkgDisplayName := bolden(pkgName)

		if !pkgExists(pkgName) {
			log(3, "%s is not installed, so it can't be updated.", pkgDisplayName)

			continue
		}

		downloadPkg(pkgName)

		chapLog("==>", "GREEN", "Success")
		log(0, "Successfully updated package info for %s.", pkgDisplayName)
	}

	chapLog("=>", "GREEN", "Success")
	log(0, "Successfully updated info for all selected packages.")
}

func updateAllPackages() {
	fullInit()

	chapLog("==>", "", "Getting installed packages")
	installedPackages := make([]string, 0)

	log(1, "Getting contents of %s", bolden(infoPath))
	files := dirContents(infoPath, "An error occurred while getting list of installed packages")

	log(1, "Iterating through contents of %s", bolden(infoPath))
	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}

	chapLog("=>", "", "Updating packages")
	for _, installedPackage := range installedPackages {
		chapLog("==>", "", "Updating %s", installedPackage)
		downloadPkg(installedPackage)

		chapLog("===>", "GREEN", "Success")
		log(0, "Successfully updated package info for %s.", bolden(installedPackage))
	}

	chapLog("=>", "", "Success")
	log(0, "Successfully updated info for all packages.")
}

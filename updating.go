package main

import (
	"strings"
)

func rawGetInfo(pkgName string, pkg Package) {
	log(1, "Getting info for %s", bolden(pkgName))
	log(1, "Checking for info URL...")

	if pkg.InfoURL == "" {
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

		rawGetInfo(installedPackage, readLoad(installedPackage))

		chapLog("===>", textCol.Green, "Success")
		log(0, "Successfully updated package info for %s.", bolden(installedPackage))
	}

	chapLog("=>", textCol.Green, "Success")
	log(0, "Successfully updated info for all packages.")
}

package main

import (
	"strings"
)

func upgradePackage(pkgNames []string) {
	loadConfig()

	for _, pkgName := range pkgNames {
		pkgDisplayName := bolden(pkgName)

		if !pkgExists(pkgName) {
			if force {
				log(3, "%s is not installed, but force is on, so continuing.", pkgDisplayName)
			} else {
				log(3, "%s is not installed, so it can't be upgraded.", pkgDisplayName)
				continue
			}
		}

		log(1, "Updating source code for %s...", pkgDisplayName)
		err := pullRepo(pkgName)

		if err.Error() == "already up-to-date" {
			continue
		}

		pkg := readLoad(pkgName)
		cmds := getUpdCmd(pkg)

		runCmds(cmds, pkg, srcPath+pkg.Name, "upgrade")
		copyBins(pkg)

		log(0, "Successfully upgraded %s!\n", pkgName)
	}
}

func upgradeAllPackages() {
	loadConfig()

	var installedPackages []string
	files := dirContents(infoPath, "An error occurred while getting list of installed packages")

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}

	log(1, "Upgrading all packages...")
	for _, installedPackage := range installedPackages {
		installedPackageDisplay := bolden(installedPackage)
		err := pullRepo(installedPackage)

		if err.Error() == "already up-to-date" {
			continue
		}

		log(1, "Upgrading %s...", installedPackageDisplay)

		pkg := readLoad(installedPackage)
		cmds := getUpdCmd(pkg)

		runCmds(cmds, pkg, srcPath+pkg.Name, "upgrade")
		copyBins(pkg)
	}

	log(0, "Upgraded all packages!")
}

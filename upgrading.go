package main

import (
	"strings"
)

func upgradePackage(pkgNames []string) {
	chapLog("=>", "VIOLET", "Initializing")
	loadConfig()

	for _, pkgName := range pkgNames {
		chapLog("=>", "VIOLET", "Upgrading %s", pkgName)
		pkgDisplayName := bolden(pkgName)

		chapLog("==>", "BLUE", "Running checks")
		if !pkgExists(pkgName) {
			if force {
				log(3, "%s is not installed, but force is on, so continuing.", pkgDisplayName)
			} else {
				log(3, "%s is not installed, so it can't be upgraded.", pkgDisplayName)
				continue
			}
		}

		chapLog("==>", "BLUE", "Pulling source code")
		log(1, "Updating source code for %s...", pkgDisplayName)
		err := pullRepo(pkgName)

		if err.Error() == "already up-to-date" {
			continue
		}

		chapLog("==>", "BLUE", "Upgrade info")
		pkg := readLoad(pkgName)
		cmds := getUpdCmd(pkg)

		if len(cmds) > 0 {
			chapLog("==>", "BLUE", "Compiling")
			runCmds(cmds, pkg, srcPath+pkg.Name, "upgrade")
		}

		chapLog("==>", "BLUE", "Installing")
		copyBins(pkg)

		chapLog("=>", "GREEN", "Success")
		log(0, "Successfully upgraded %s!", pkgName)
	}
}

func upgradeAllPackages() {
	chapLog("=>", "VIOLET", "Initializing")
	loadConfig()

	chapLog("==>", "BLUE", "Getting installed packages")
	var installedPackages []string
	files := dirContents(infoPath, "An error occurred while getting list of installed packages")

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}

	chapLog("=>", "VIOLET", "Starting upgrades")
	for _, installedPackage := range installedPackages {
		chapLog("==>", "BLUE", "Upgrading %s", installedPackage)
		err := pullRepo(installedPackage)

		if err.Error() == "already up-to-date" {
			continue
		}

		pkg := readLoad(installedPackage)
		cmds := getUpdCmd(pkg)

		runCmds(cmds, pkg, srcPath+pkg.Name, "upgrade")
		copyBins(pkg)
	}

	chapLog("=>", "GREEN", "Success")
	log(0, "Upgraded all packages!")
}

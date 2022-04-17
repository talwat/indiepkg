package main

import (
	"strings"
)

func upgradePackage(pkgNames []string) {
	fullInit()

	for _, pkgName := range pkgNames {
		chapLog("=>", "VIOLET", "Upgrading %s", pkgName)
		pkgDisplayName := bolden(pkgName)

		chapLog("==>", "BLUE", "Running checks")
		log(1, "Checking if %s exists...", pkgDisplayName)
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

		directDownload := false

		if err.Error() == "already up-to-date" {
			if force {
				log(3, "%s is already up to date, but force is on, so continuing.", bolden(pkgName))
			} else {
				log(0, "%s already up to date.", bolden(pkgName))
				continue
			}
		} else if err.Error() == "repository does not exist" && pathExists(srcPath+pkgName, "An error occurred while checking if %s's source exists", pkgName) {
			log(1, "Direct download detected.")
			directDownload = true
		}

		chapLog("==>", "BLUE", "Upgrade info")
		pkg := readLoad(pkgName)
		cmds := getUpdCmd(pkg)

		if directDownload {
			chapLog("==>", "BLUE", "Updating info")
			oldVer := pkg.Version
			log(1, "Getting & writing new info for %s...", pkgDisplayName)
			writeLoadPkg(pkgName, findPkg(pkgName), false)
			log(1, "Reading new version number...")
			newVer := readLoad(pkgName).Version

			debugLog("Old version: %s. New version: %s", oldVer, newVer)
			chapLog("==>", "BLUE", "Checking if already up to date")
			log(1, "Checking if %s is already up to date...", bolden(pkgName))
			if oldVer == newVer {
				log(0, "%s already up to date.", pkgDisplayName)
				continue
			} else {
				log(1, "Not up to date. Upgrading from %s to %s", bolden(oldVer), bolden(newVer))
			}

			doDirectDownload(pkg, pkgName, srcPath)
		}

		if len(cmds) > 0 {
			chapLog("==>", "BLUE", "Compiling")
			runCmds(cmds, pkg, srcPath+pkg.Name, "upgrade")
		}

		chapLog("==>", "BLUE", "Installing")
		copyBins(pkg, srcPath)

		chapLog("=>", "GREEN", "Success")
		log(0, "Successfully upgraded %s!", pkgName)
	}
}

func upgradeAllPackages() {
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
		err := pullRepo(installedPackage)

		if err.Error() == "already up-to-date" {
			continue
		}

		pkg := readLoad(installedPackage)
		cmds := getUpdCmd(pkg)

		runCmds(cmds, pkg, srcPath+pkg.Name, "upgrade")
		copyBins(pkg, srcPath)
	}

	chapLog("=>", "GREEN", "Success")
	log(0, "Upgraded all packages!")
}

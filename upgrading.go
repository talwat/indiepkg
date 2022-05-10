package main

import (
	"strings"
)

func upgradePkgFunc(pkgName string, chapPrefix string) {
	chapLog(chapPrefix+"=>", "", "Upgrading %s", pkgName)

	pkgDisplayName := bolden(pkgName)

	chapLog(chapPrefix+"==>", "", "Running checks")
	log(1, "Checking if %s exists...", pkgDisplayName)

	if !pkgExists(pkgName) {
		if force {
			log(3, "%s is not installed, but force is on, so continuing.", pkgDisplayName)
		} else {
			log(3, "%s is not installed, so it can't be upgraded.", pkgDisplayName)

			return
		}
	}

	chapLog(chapPrefix+"==>", "", "Pulling source code")
	log(1, "Updating source code for %s...", pkgDisplayName)

	isUpToDate, directDownload := pullPkgRepo(pkgName)

	log(1, "Checking if up to date...")

	if isUpToDate {
		if force {
			log(3, "%s already up to date, but force is on, so continuing.", pkgDisplayName)
		} else {
			log(0, "%s already up to date.", pkgDisplayName)

			return
		}
	}

	chapLog(chapPrefix+"==>", "", "Reading package info")

	pkg := readLoad(pkgName)

	if directDownload {
		chapLog(chapPrefix+"==>", "", "Updating info")

		log(1, "Getting & writing new info for %s...", pkgDisplayName)
		log(1, "Checking for info URL...")

		rawGetInfo(pkgName, pkg)

		chapLog(chapPrefix+"==>", "", "Getting version numbers")
		log(1, "Reading new version number...")

		newVer := readLoad(pkgName).Version

		rawLog("\n")

		log(1, "Saving old version number...")

		oldVer := pkg.Version

		debugLog("Old version: %s. New version: %s", oldVer, newVer)

		chapLog(chapPrefix+"==>", "", "Checking if already up to date")
		log(1, "Checking if %s is already up to date...", bolden(pkgName))

		if oldVer == newVer {
			if !force {
				log(0, "%s already up to date.", pkgDisplayName)

				return
			}

			log(3, "%s already up to date, but force is on, so continuing.", pkgDisplayName)
		} else {
			log(1, "Not up to date. Upgrading from %s to %s", bolden(oldVer), bolden(newVer))
		}

		chapLog(chapPrefix+"==>", "", "Downloading file")
		doDirectDownload(pkg, srcPath)
	}

	chapLog(chapPrefix+"==>", "", "Getting upgrade commands")

	if cmds := getUpdCmd(pkg); len(cmds) > 0 {
		chapLog(chapPrefix+"==>", "", "Compiling")
		runCmds(cmds, pkg, srcPath+pkg.Name, "upgrade")
	}

	chapLog(chapPrefix+"==>", "", "Upgrading")
	copyBins(pkg, srcPath)
	copyManpages(pkg, srcPath)

	chapLog(chapPrefix+"==>", "GREEN", "Success")
	log(0, "Successfully upgraded %s!", pkgDisplayName)
}

func upgradePackage(pkgNames []string) {
	fullInit()

	for _, pkgName := range pkgNames {
		upgradePkgFunc(pkgName, "")
	}
}

func upgradeAllPackages() {
	fullInit()

	chapLog("==>", "", "Getting installed packages")

	installedPackages := make([]string, 0)

	files := dirContents(infoPath, "An error occurred while getting list of installed packages")

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}

	chapLog("=>", "", "Starting upgrades")

	for _, installedPackage := range installedPackages {
		upgradePkgFunc(installedPackage, "=")
	}

	chapLog("=>", "GREEN", "Success")
	log(0, "Upgraded all packages.")
}

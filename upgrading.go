package main

import (
	"strings"
)

func upgradePackage(pkgNames []string) {
	for _, pkgName := range pkgNames {
		pkgDisplayName := bolden(pkgName)

		if pkgExists(pkgName) {
			log(3, "%s is not installed, so it can't be upgraded.", pkgDisplayName)
			continue
		}

		log(1, "Updating source code for %s...", pkgDisplayName)
		pullOutput, _ := runCommand(srcPath+pkgName, "git", "pull")

		if strings.Contains(pullOutput, "Already up to date") {
			log(0, "%s already up to date.", pkgDisplayName)
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
	var installedPackages []string
	files := dirContents(installedPath, "An error occurred while getting list of installed packages")

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}

	log(1, "Upgrading all packages...")
	for _, installedPackage := range installedPackages {
		installedPackageDisplay := bolden(installedPackage)
		pullOutput, _ := runCommand(srcPath+installedPackage, "git", "pull")

		if strings.Contains(pullOutput, "Already up to date") {
			log(0, "%s already up to date.", installedPackageDisplay)
			continue
		}

		log(1, "Upgrading %s...", installedPackageDisplay)

		pkg := readLoad(installedPackage)

		cmds := getUpdCmd(pkg)

		runCmds(cmds, pkg, srcPath+pkg.Name, "upgrade")

		copyBins(pkg)

		runCmds(getUpdCmd(pkg), pkg, srcPath+pkg.Name, "upgrade")
	}

	log(0, "Upgraded all packages!")
}

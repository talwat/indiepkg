package main

import (
	"strings"
)

func upgradePackage(pkgNames []string) {
	for _, pkgName := range pkgNames {
		pkgSrcPath := home + "/.local/share/indiepkg/package_src"

		if !packageExists(pkgName) {
			log(3, "%s is not installed, so it can't be upgraded.", pkgName)
			continue
		}

		pkg := readAndLoad(pkgName)

		log(1, "Updating source code...")
		runCommand(pkgSrcPath+"/"+pkgName, "git", "pull")

		log(1, "Running upgrade commands...")
		runCommands(pkg.Update, pkg)

		log(0, "Successfully upgraded %s!\n", pkgName)
	}
}

func upgradeAllPackages() {
	srcPath := home + "/.local/share/indiepkg/package_src/"
	infoPath := home + "/.local/share/indiepkg/installed_packages/"
	var installedPackages []string
	files := dirContents(infoPath, "An error occurred while getting list of installed packages")

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}
	log(1, "Updating all packages...")
	for _, installedPackage := range installedPackages {
		pullOutput, _ := runCommand(srcPath+installedPackage, "git", "pull")
		if strings.Contains(pullOutput, "Already up to date") {
			log(0, "%s already up to date.", installedPackage)
			continue
		}
		log(1, "Updating %s", installedPackage)

		pkg := readAndLoad(installedPackage)

		runCommands(pkg.Update, pkg)
	}

	log(0, "Upgraded all packages!")
}

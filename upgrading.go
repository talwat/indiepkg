package main

import (
	"os"
	"strings"
)

func upgradePackage(pkgName string) {
	pkgSrcPath := home + "/.local/share/indiepkg/package_src"
	pkgInfoPath := home + "/.local/share/indiepkg/installed_packages/" + pkgName + ".json"
	var err error

	installed, err := pathExists(pkgInfoPath)
	errorLog(err, 4, "An error occurred while checking if package %s exists.", pkgName)
	if !installed {
		log(4, "%s is not installed, so it can't be upgraded.", pkgName)
		os.Exit(1)
	}

	log(1, "Reading package info...")
	pkgFile, err := readFile(pkgInfoPath)
	errorLog(err, 4, "An error occurred while reading package information for %s.", pkgName)

	log(1, "Loading package info...")
	pkg, err := loadPackage(pkgFile)
	errorLog(err, 4, "An error occurred while loading package information for %s.", pkgName)

	log(1, "Updating source code...")
	runCommand(pkgSrcPath+"/"+pkgName, "git", "pull")

	log(1, "Running upgrade commands...")
	for _, command := range pkg.upgrade {
		runCommand(pkgSrcPath+"/"+pkg.Name, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
	}

	log(0, "Successfully upgraded %s!", pkgName)
}

func upgradeAllPackages() {
	srcPath := home + "/.local/share/indiepkg/package_src/"
	infoPath := home + "/.local/share/indiepkg/installed_packages/"
	var err error
	var installedPackages []string
	files, err := dirContents(infoPath)
	errorLog(err, 4, "An error occurred while getting list of installed packages.")

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}
	log(1, "Updating all packages...")
	for _, installedPackage := range installedPackages {
		pullOutput, _ := runCommand(srcPath+installedPackage, "git", "pull")
		if strings.Contains(pullOutput, "Already up to date") {
			continue
		}
		log(1, "Updating %s", installedPackage)
		errorLog(err, 4, "An error occurred while getting package information for %s.", installedPackage)
		pkgFile, err := readFile(infoPath + installedPackage)
		errorLog(err, 4, "An error occurred while reading package information for %s.", installedPackage)
		pkg, _ := loadPackage(pkgFile)
		for _, command := range pkg.upgrade {
			runCommand(srcPath+installedPackage+"/"+pkg.Name, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
		}
	}

	log(0, "Upgraded all packages!")
}

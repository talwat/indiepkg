package main

import (
	"os"
	"strings"
)

func updatePackage(pkgName string) {
	pkgInfoPath := home + "/.local/share/indiepkg/installed_packages/" + pkgName + ".json"
	url := "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json"
	var err error

	installed, err := pathExists(pkgInfoPath)
	errorLog(err, 4, "An error occurred while checking if package %s exists.", pkgName)
	if !installed {
		log(4, "%s is not installed, so it can't be updated.", pkgName)
		os.Exit(1)
	}

	log(1, "Downloading package info...")
	log(1, "URL: %s", url)
	err = downloadFile(pkgInfoPath, url)
	errorLog(err, 4, "An error occurred while getting package information for %s.", pkgName)

	log(0, "Successfully updated package info for %s!", pkgName)
}

func updateAllPackages() {
	pkgInfoPath := home + "/.local/share/indiepkg/installed_packages/"
	url := "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/"
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
		err = downloadFile(pkgInfoPath+installedPackage+".json", url+installedPackage+".json")
		errorLog(err, 4, "An error occurred while getting package information for %s.", installedPackage)
	}

	log(0, "Successfully updated info for all packages!")
}

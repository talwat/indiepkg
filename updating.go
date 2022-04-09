package main

import (
	"os"
	"strings"
)

func updatePackage(pkgName string) {
	url := "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json"

	if !packageExists(pkgName) {
		log(4, "%s is not installed, so it can't be updated.", pkgName)
		os.Exit(1)
	}

	log(1, "Downloading package info...")
	log(1, "URL: %s", url)
	downloadFile(installedPath, url, "An error occurred while getting package information for %s", pkgName)

	log(0, "Successfully updated package info for %s!", pkgName)
}

func updateAllPackages() {
	url := "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/"

	var installedPackages []string
	files := dirContents(installedPath, "An error occurred while getting list of installed packages")

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}
	log(1, "Updating all packages...")
	for _, installedPackage := range installedPackages {
		downloadFile(installedPath+installedPackage+".json", url+installedPackage+".json", "An error occurred while getting package information for %s", installedPackage)
	}

	log(0, "Successfully updated info for all packages!")
}

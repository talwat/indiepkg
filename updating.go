package main

import (
	"strings"
)

func updatePackage(pkgNames []string) {
	for _, pkgName := range pkgNames {
		pkgDisplayName := bolden(pkgName)

		url := "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json"

		if !packageExists(pkgName) {
			log(3, "%s is not installed, so it can't be updated.", pkgDisplayName)
			continue
		}

		log(1, "Downloading package info for %s...", pkgDisplayName)
		log(1, "URL: %s", url)
		downloadFile(installedPath+pkgName+".json", url, "An error occurred while getting package information for %s", pkgDisplayName)

		log(0, "Successfully updated package info for %s!\n", pkgDisplayName)
	}
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
		log(1, "Downloading package info for %s...", bolden(installedPackage))
		downloadFile(installedPath+installedPackage+".json", url+installedPackage+".json", "An error occurred while getting package information for %s", bolden(installedPackage))
	}

	log(0, "Successfully updated info for all packages!")
}

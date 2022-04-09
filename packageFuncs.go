package main

import "fmt"

func readAndLoad(packageName string) Package {
	log(1, "Reading package info for %s...", packageName)
	pkgFile := readFile(installedPath+packageName+".json", "An error occurred while reading package %s", packageName)

	log(1, "Loading package info for %s...", packageName)
	pkg := loadPackage(pkgFile, fmt.Sprintf("An error occurred while loading package information for %s", packageName))

	return pkg
}

func packageExists(pkgName string) bool {
	infoInstalled := pathExists(installedPath+pkgName+".json", "An error occurred while checking if package info for %s exists", pkgName)

	srcInstalled := pathExists(srcPath+pkgName, "An error occurred while checking if package source for %s exists", pkgName)

	if infoInstalled && srcInstalled {
		return true
	} else if !infoInstalled && !srcInstalled {
		return false
	} else {
		log(4, "Package info or source for %s exists, but not both. Please run %sindiepkg repair%s", pkgName, textFx["BOLD"], RESETCOL)
		return false
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func listPkgs() {
	var installedPkgs []string
	files := dirContents(infoPath, "An error occurred while getting list of installed packages")

	if len(files) == 0 {
		log(1, "No packages installed.")
		os.Exit(0)
	}

	for _, file := range files {
		installedPkgs = append(installedPkgs, strings.ReplaceAll(file.Name(), ".json", ""))
	}
	fmt.Println(strings.Join(installedPkgs, "\n"))
}

func sync() {
	fullInit()

	chapLog("==>", "BLUE", "Getting packages to sync")

	log(1, "Getting list of package files...")
	dirs := dirContents(srcPath, "An error occurred while getting list of source files")

	log(1, "Getting missing package info...")
	var pkgInfoToSync []string
	for _, dir := range dirs {
		pkgName := strings.ReplaceAll(dir.Name(), ".json", "")
		infoExists := pathExists(infoPath+pkgName+".json", "An error occurred while checking if %s is properly installed", pkgName)
		if !infoExists && dir.IsDir() {
			pkgInfoToSync = append(pkgInfoToSync, pkgName)
		}
	}

	log(1, "Getting list of source directories...")
	infoFiles := dirContents(infoPath, "An error occurred while getting list of info files")

	log(1, "Getting missing source directories...")
	var pkgSrcToSync []string
	for _, infoFile := range infoFiles {
		pkgName := strings.ReplaceAll(infoFile.Name(), ".json", "")
		srcExists := pathExists(srcPath+pkgName, "An error occurred while checking if %s is properly installed", pkgName)
		if !srcExists && !infoFile.IsDir() {
			pkgSrcToSync = append(pkgSrcToSync, pkgName)
		}
	}

	chapLog("=>", "VIOLET", "Syncing packages...")
	for _, pkgToSync := range pkgInfoToSync {
		chapLog("==>", "BLUE", "Downloading info for %s", pkgToSync)
		downloadPkg(pkgToSync, false)
	}

	for _, pkgToSync := range pkgSrcToSync {
		chapLog("==>", "BLUE", "Cloning source for %s", pkgToSync)
		cloneRepo(readLoad(pkgToSync), srcPath)
	}

	if len(pkgInfoToSync) > 0 || len(pkgSrcToSync) > 0 {
		chapLog("=>", "GREEN", "Success")
		log(0, "Successfully synced info for %s and source for %s!", strings.Join(pkgInfoToSync, ", "), strings.Join(pkgSrcToSync, ", "))
	} else {
		chapLog("=>", "CYAN", "Nothing synced")
		log(1, "Nothing synced.")
	}
}

func infoPkg(pkgName string) {
	pkg, _ := getPkgFromNet(pkgName)
	fmt.Printf("\n")
	log(1, "Name: %s", pkg.Name)
	log(1, "Author: %s", pkg.Author)
	log(1, "Description: %s", pkg.Description)
	log(1, "License: %s", pkg.License)
	log(1, "Git URL: %s", pkg.Url)

	deps := getDeps(pkg)
	if deps != nil {
		log(1, "Dependencies: %s", strings.Join(deps, ", "))
	}

	getNotes(pkg)
}

func rmData(pkgNames []string) {
	log(3, "Warning: This will remove the data for the selected packages stored in %s", mainPath)
	log(3, "This will %snot%s run the uninstall commands.", textFx["BOLD"], RESETCOL)
	log(3, "You should only use this in case a package installation has failed at a certain step, or you want to separate an installed package from indiepkg.")
	displayPkgs(pkgNames, "remove the data for")

	for _, pkgName := range pkgNames {
		chapLog("=>", "VIOLET", "Removing data for %s", pkgName)
		pkgDisplayName := bolden(pkgName)

		log(1, "Deleting source files for %s...", pkgDisplayName)
		delPath(3, srcPath+pkgName, "An error occurred while deleting source files for %s", pkgDisplayName)

		log(1, "Deleting info file for %s...", pkgDisplayName)
		delPath(3, infoPath+pkgName+".json", "An error occurred while deleting info file for %s", pkgDisplayName)

		log(0, "Successfully deleted the data for %s.\n", pkgDisplayName)
	}

	chapLog("=>", "GREEN", "Success")
	log(0, "Successfully deleted data.")
}

func search(query string) {
	pkgs := getPkgFromGh(query)

	log(1, "Found %d packages:", len(pkgs))
	for _, pkg := range pkgs {
		fmt.Println("        " + pkg.Name + " - " + pkg.Repo)
	}
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func listPackages() {
	var installedPackages []string
	files := dirContents(installedPath, "An error occurred while getting list of installed packages")

	if len(files) == 0 {
		log(1, "No packages installed.")
		os.Exit(0)
	}

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}
	fmt.Println(strings.Join(installedPackages, "\n"))
}

func repair() {
	initDirs("Making required directories...")

	dirs := dirContents(srcPath, "An error occurred while getting list of source files")

	var packageInfoToRepair []string
	for _, dir := range dirs {
		packageName := strings.ReplaceAll(dir.Name(), ".json", "")
		infoExists := pathExists(installedPath+packageName+".json", "An error occurred while checking if %s is properly installed", packageName)
		if !infoExists && dir.IsDir() {
			packageInfoToRepair = append(packageInfoToRepair, packageName)
		}
	}

	for _, packageToRepair := range packageInfoToRepair {
		downloadFile(installedPath+packageToRepair+".json", "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/"+packageToRepair+".json", "An error occurred while downloading package information for %s", packageToRepair)
	}

	infoFiles := dirContents(installedPath, "An error occurred while getting list of info files")

	var packageSrcToRepair []string
	for _, infoFile := range infoFiles {
		packageName := strings.ReplaceAll(infoFile.Name(), ".json", "")
		srcExists := pathExists(srcPath+packageName, "An error occurred while checking if %s is properly installed", packageName)
		if !srcExists && !infoFile.IsDir() {
			packageSrcToRepair = append(packageSrcToRepair, packageName)
		}
	}

	for _, packageToRepair := range packageInfoToRepair {
		log(1, "Downloading package info for %s...", packageToRepair)
		downloadFile(installedPath+packageToRepair+".json", "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/"+packageToRepair+".json", "An error occurred while downloading package information for %s", packageToRepair)
	}

	for _, packageToRepair := range packageSrcToRepair {
		log(1, "Repairing package %s...", packageToRepair)

		pkg := readAndLoad(packageToRepair)

		log(1, "Cloning package source for %s...", packageToRepair)
		output, exit_code := runCommand(srcPath, "git", "clone", pkg.Url)
		log(1, output)
		if exit_code != 0 {
			errorLog(errors.New(fmt.Sprintf("Command exited with code %d", exit_code)), 4, "An error occurred while cloning package source for %s", packageToRepair)
		}
	}

	if len(packageInfoToRepair) > 0 || len(packageSrcToRepair) > 0 {
		log(0, "Successfully repaired %s!", strings.Join(packageInfoToRepair, ", ")+", "+strings.Join(packageSrcToRepair, ", "))
	} else {
		log(1, "Nothing repaired.")
	}
}

func infoPackage(pkgName string) {
	pkg, _ := getPkgFromNet(pkgName)
	log(1, "Name: %s", pkg.Name)
	log(1, "Author: %s", pkg.Author)
	log(1, "Description: %s", pkg.Description)
	log(1, "Git URL: %s", pkg.Url)

	deps := getDeps(pkg)
	if deps != nil {
		log(1, "Dependencies: %s", strings.Join(deps, ", "))
	}
}

func removeData(pkgNames []string) {
	log(3, "Warning: This will remove the data for the selected packages stored in %s", mainPath)
	log(3, "This will %snot%s run the uninstall commands.", textFx["BOLD"], RESETCOL)
	log(1, "Are you sure you would like to remove the data for the following packages:")
	for _, packageToRemove := range pkgNames {
		fmt.Println("        " + packageToRemove)
	}

	confirm("y", "(y/n)")

	for _, pkgName := range pkgNames {
		log(1, "Deleting source files for %s...", pkgName)
		delPath(3, srcPath+pkgName, "An error occurred while deleting source files for %s", pkgName)

		log(1, "Deleting info file for %s...", pkgName)
		delPath(3, installedPath+pkgName+".json", "An error occurred while deleting info file for package %s", pkgName)

		log(0, "Successfully deleted the data for %s.\n", pkgName)
	}
}

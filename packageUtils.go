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

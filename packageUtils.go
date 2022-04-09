package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func loadPackage(packageFile string) (Package, error) {
	var pkg Package

	keySlice := make([]string, 0)
	for key := range environmentVariables {
		keySlice = append(keySlice, key)
	}

	for _, key := range keySlice {
		packageFile = strings.Replace(packageFile, ":("+key+"):", environmentVariables[key], -1)
	}
	err := json.Unmarshal([]byte(packageFile), &pkg)
	return pkg, err
}

func listPackages() {
	var err error
	var installedPackages []string
	files, err := dirContents(home + "/.local/share/indiepkg/installed_packages/")
	errorLog(err, 4, "An error occurred while getting list of installed packages.")

	for _, file := range files {
		installedPackages = append(installedPackages, strings.ReplaceAll(file.Name(), ".json", ""))
	}
	fmt.Println(strings.Join(installedPackages, "\n"))
}

func repair() {
	log(1, "Making required directories...")
	newDir(installedPath) //nolint:errcheck
	newDir(srcPath)       //nolint:errcheck

	dirs, err := dirContents(srcPath)
	errorLog(err, 4, "An error occurred while getting list of source files.")

	var packageInfoToRepair []string
	for _, dir := range dirs {
		packageName := strings.ReplaceAll(dir.Name(), ".json", "")
		infoExists, err := pathExists(installedPath + packageName + ".json")
		errorLog(err, 4, "An error occurred while checking if %s is properly installed.", packageName)
		if !infoExists && dir.IsDir() {
			packageInfoToRepair = append(packageInfoToRepair, packageName)
		}
	}

	for _, packageToRepair := range packageInfoToRepair {
		err = downloadFile(installedPath+packageToRepair+".json", "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/"+packageToRepair+".json")
		errorLog(err, 4, "An error occurred while downloading package information for %s.", packageToRepair)
	}

	infoFiles, err := dirContents(installedPath)
	errorLog(err, 4, "An error occurred while getting list of info files.")

	var packageSrcToRepair []string
	for _, infoFile := range infoFiles {
		packageName := strings.ReplaceAll(infoFile.Name(), ".json", "")
		srcExists, err := pathExists(srcPath + packageName)
		errorLog(err, 4, "An error occurred while checking if %s is properly installed.", packageName)
		if !srcExists && !infoFile.IsDir() {
			packageSrcToRepair = append(packageSrcToRepair, packageName)
		}
	}

	for _, packageToRepair := range packageInfoToRepair {
		log(1, "Downloading package info for %s...", packageToRepair)
		err = downloadFile(installedPath+packageToRepair+".json", "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/"+packageToRepair+".json")
		errorLog(err, 4, "An error occurred while downloading package information for %s.", packageToRepair)
	}

	for _, packageToRepair := range packageSrcToRepair {
		log(1, "Repairing package %s...", packageToRepair)

		pkg := readAndLoad(packageToRepair)

		log(1, "Cloning package source for %s...", packageToRepair)
		output, exit_code := runCommand(srcPath, "git", "clone", pkg.Url)
		log(1, output)
		if exit_code != 0 {
			errorLog(err, 4, "An error occurred while cloning package source for %s.", packageToRepair)
		}
	}

	if len(packageInfoToRepair) > 0 || len(packageSrcToRepair) > 0 {
		log(0, "Successfully repaired %s!", strings.Join(packageInfoToRepair, ", ")+", "+strings.Join(packageSrcToRepair, ", "))
	} else {
		log(1, "Nothing repaired.")
	}
}

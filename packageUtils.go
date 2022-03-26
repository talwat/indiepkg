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

	var packagesToRepair []string
	for _, dir := range dirs {
		packageName := strings.ReplaceAll(dir.Name(), ".json", "")
		infoExists, err := pathExists(installedPath + packageName + ".json")
		errorLog(err, 4, "An error occurred while checking if %s is properly installed.", packageName)
		if !infoExists && dir.IsDir() {
			packagesToRepair = append(packagesToRepair, packageName)
		}
	}

	for _, packageToRepair := range packagesToRepair {
		err = downloadFile(installedPath+packageToRepair+".json", "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/"+packageToRepair+".json")
		errorLog(err, 4, "An error occurred while downloading package information for %s.", packageToRepair)
	}

	if len(packagesToRepair) > 0 {
		log(0, "Successfully repaired %s!", strings.Join(packagesToRepair, ", "))
	} else {
		log(1, "Nothing repaired.")
	}
}

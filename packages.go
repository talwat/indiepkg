package main

import (
	"encoding/json"
	"os"
	"strings"
)

var home string = os.Getenv("HOME")

type Package struct {
	Name         string
	Author       string
	Description  string
	Url          string
	Install      []string
	Uninstall    []string
	Update       []string
	Config_paths []string
}

var environmentVariables = map[string]string{
	"PATH": home + "/.local",
}

func loadPackage(packageFile string, pkgName string) (Package, error) {
	var pkg Package

	log(1, "Loading package info...")
	keySlice := make([]string, 0)
	for key := range environmentVariables {
		keySlice = append(keySlice, key)
	}

	for _, key := range keySlice {
		packageFile = strings.Replace(packageFile, ":("+key+"):", environmentVariables[key], -1)
	}
	err := json.Unmarshal([]byte(packageFile), &pkg)
	if err != nil {
		errorLog(err, 4, "An error occurred while loading package %s.", pkgName)
	}
	return pkg, err
}

func installPackage(pkgName string) {
	pkgSrcPath := home + "/.local/share/indiepkg/package_src"
	pkgInfoPath := home + "/.local/share/indiepkg/installed_packages/" + pkgName + ".json"
	installedPkgsPath := home + "/.local/share/indiepkg/installed_packages/"
	url := "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json"
	var err error

	log(1, "Making required directories...")
	newDir(pkgSrcPath)        //nolint:errcheck
	newDir(installedPkgsPath) //nolint:errcheck

	log(1, "Downloading package info...")
	log(1, "URL: %s", url)
	err = downloadFile(pkgInfoPath, url)
	errorLog(err, 4, "An error occurred while getting package information for %s.", pkgName)

	log(1, "Reading package info...")
	pkgFile, err := readFile(pkgInfoPath)
	errorLog(err, 4, "An error occurred while reading package information for %s.", pkgName)

	pkg, _ := loadPackage(pkgFile, pkgName)

	log(1, "Cloning source code...")
	runCommand(pkgSrcPath, "git", "clone", pkg.Url)

	log(1, "Running install commands...")
	for _, command := range pkg.Install {
		runCommand(pkgSrcPath+"/"+pkg.Name, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
	}

	log(0, "Installed %s successfully!", pkgName)
}

func uninstallPackage(pkgName string) {
	pkgSrcPath := home + "/.local/share/indiepkg/package_src/"
	pkgInfoPath := home + "/.local/share/indiepkg/installed_packages/" + pkgName + ".json"
	var err error

	log(1, "Reading package...")
	rawPkgInfo, err := readFile(pkgInfoPath)
	errorLog(err, 4, "An error occurred while reading package %s.", pkgName)

	pkg, _ := loadPackage(rawPkgInfo, pkgName)

	log(1, "Running uninstall commands...")
	for _, command := range pkg.Uninstall {
		runCommand(pkgSrcPath+"/"+pkg.Name, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
	}

	log(1, "Deleting source files for %s...", pkgName)
	err = delDir(pkgSrcPath + pkgName)
	errorLog(err, 4, "An error occurred while deleting source files for %s.", pkgName)

	log(1, "Deleting info file for %s...", pkgName)
	err = delFile(pkgInfoPath)
	errorLog(err, 4, "An error occurred while deleting info file for package %s.", pkgName)

	log(0, "Successfully uninstalled %s.", pkgName)
}

func infoPackage(pkgName string) {
	packageFile, err := viewFile("https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json")
	errorLog(err, 4, "An error occurred while getting package info for %s.", pkgName)
	pkgInfo, _ := loadPackage(packageFile, pkgName)
	log(1, "Name: %s", pkgInfo.Name)
	log(1, "Author: %s", pkgInfo.Author)
	log(1, "Description: %s", pkgInfo.Description)
	log(1, "Git URL: %s", pkgInfo.Url)
}

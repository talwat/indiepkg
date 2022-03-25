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

func loadPackage(packageFile string) (Package, error) {
	var pkg Package
	//packageFile, _ := viewFile("https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json")

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

func installPackage(pkgName string) {
	pkgSrcPath := home + "/.local/share/indiepkg/package_src"
	pkgInfoPath := home + "/.local/share/indiepkg/installed_packages/" + pkgName + ".json"
	var err error

	log(1, "Making required directories...")
	err = newDir(pkgSrcPath)
	err = newDir(home + "/.local/share/indiepkg/installed_packages/")

	log(1, "Downloading package info...")
	err = downloadFile(pkgInfoPath, "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/"+pkgName+".json")
	errorLog(err, 4, "An error occurred while getting package information for %s.", pkgName)

	log(1, "Reading package info...")
	pkgFile, err := readFile(pkgInfoPath)
	errorLog(err, 4, "An error occurred while reading package information for %s.", pkgName)

	log(1, "Loading package...")
	pkg, err := loadPackage(pkgFile)
	errorLog(err, 4, "An error occurred while loading package %s.", pkgName)

	log(1, "Cloning source code...")
	runCommand(pkgSrcPath, "git", "clone", pkg.Url)

	log(1, "Running install commands...")
	for _, command := range pkg.Install {
		runCommand(pkgSrcPath+"/"+pkg.Name, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
	}
}

func uninstallPackage(pkgName string) {
	pkgSrcPath := home + "/.local/share/indiepkg/package_src"
	pkgInfoPath := home + "/.local/share/indiepkg/installed_packages/" + pkgName + ".json"

	log(1, "Reading package...")
	rawPkgInfo, err := readFile(pkgInfoPath)
	errorLog(err, 4, "An error occurred while reading package %s.", pkgName)

	log(1, "Loading package...")
	pkg, err := loadPackage(rawPkgInfo)
	errorLog(err, 4, "An error occurred while loading package %s.", pkgName)
	log(1, "Running uninstall commands...")
	for _, command := range pkg.Uninstall {
		runCommand(pkgSrcPath+"/"+pkg.Name, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
	}
	log(1, "Deleting source files...")
	delDir(home + "/.local/share/indiepkg/package_src/" + pkgName)
	log(0, "Successfully uninstalled ")
}

func infoPackage(pkgName string) {
	packageFile, err := viewFile("https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json")
	errorLog(err, 4, "An error occurred while getting package info for %s.", pkgName)
	pkgInfo, err := loadPackage(packageFile)
	errorLog(err, 4, "An error occurred while loading package %s.", pkgName)
	log(1, "Name: %s", pkgInfo.Name)
	log(1, "Author: %s", pkgInfo.Author)
	log(1, "Description: %s", pkgInfo.Description)
	log(1, "Git URL: %s", pkgInfo.Url)
}

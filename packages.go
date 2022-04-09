package main

import (
	"fmt"
	"os"
	"strings"
)

var home string = os.Getenv("HOME")

var srcPath string = home + "/.local/share/indiepkg/package_src/"
var installedPath string = home + "/.local/share/indiepkg/installed_packages/"

type Package struct {
	Name         string
	Author       string
	Description  string
	Url          string
	Install      []string
	Uninstall    []string
	upgrade      []string
	Config_paths []string
}

var environmentVariables = map[string]string{
	"PATH": home + "/.local",
}

func installPackages(pkgNames []string) {
	log(1, "Are you sure you would like to install the following packages:")
	for _, packageToInstall := range pkgNames {
		fmt.Println("        " + packageToInstall)
	}

	confirm("(y/n)")

	for _, pkgName := range pkgNames {
		pkgInfoPath := installedPath + pkgName + ".json"
		url := "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json"
		var err error
		log(1, "Getting package info for %s...", pkgName)
		log(1, "URL: %s", url)
		pkgInfo, err := viewFile(url)
		errorLog(err, 4, "An error occurred while getting package information for %s.", pkgName)

		log(1, "Making required directories for %s...", pkgName)
		newDir(srcPath)       //nolint:errcheck
		newDir(installedPath) //nolint:errcheck

		log(1, "Writing package info for... %s", pkgName)
		err = newFile(pkgInfoPath, pkgInfo)
		errorLog(err, 4, "An error occurred while writing package information for %s.", pkgName)

		pkg := readAndLoad(pkgName)

		log(1, "Cloning source code for %s...", pkgName)
		runCommand(srcPath, "git", "clone", pkg.Url)

		log(1, "Running install commands for %s...", pkgName)
		for _, command := range pkg.Install {
			runCommand(srcPath+pkg.Name, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
		}

		log(0, "Installed %s successfully!\n", pkgName)
	}
}

func uninstallPackages(pkgNames []string) {
	log(1, "Are you sure you would like to uninstall the following packages:")
	for _, packageToUninstall := range pkgNames {
		fmt.Println("        " + packageToUninstall)
	}

	confirm("(y/n)")

	for _, pkgName := range pkgNames {
		pkgInfoPath := installedPath + pkgName + ".json"
		var err error

		installed, err := pathExists(pkgInfoPath)
		errorLog(err, 4, "An error occurred while checking if package %s exists.", pkgName)
		if !installed {
			log(4, "%s is not installed, so it can't be uninstalled.", pkgName)
			os.Exit(1)
		}

		pkg := readAndLoad(pkgName)

		log(1, "Running uninstall commands...")
		for _, command := range pkg.Uninstall {
			runCommand(srcPath+pkg.Name, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
		}

		log(1, "Deleting source files for %s...", pkgName)
		err = delDir(srcPath + pkgName)
		errorLog(err, 4, "An error occurred while deleting source files for %s.", pkgName)

		log(1, "Deleting info file for %s...", pkgName)
		err = delFile(pkgInfoPath)
		errorLog(err, 4, "An error occurred while deleting info file for package %s.", pkgName)

		log(0, "Successfully uninstalled %s.", pkgName)
	}
}

func infoPackage(pkgName string) {
	packageFile, err := viewFile("https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json")
	errorLog(err, 4, "An error occurred while getting package info for %s.", pkgName)

	pkgInfo, err := loadPackage(packageFile)
	errorLog(err, 4, "An error occurred while loading package information for %s.", pkgName)

	log(1, "Name: %s", pkgInfo.Name)
	log(1, "Author: %s", pkgInfo.Author)
	log(1, "Description: %s", pkgInfo.Description)
	log(1, "Git URL: %s", pkgInfo.Url)
}

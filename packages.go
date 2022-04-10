package main

import (
	"fmt"
	"os"
	"strings"
)

var home string = os.Getenv("HOME") + "/"

var srcPath string = home + ".indiepkg/data/package_src/"
var installedPath string = home + ".indiepkg/data/installed_packages/"

type Deps struct {
	All     []string
	Linux   []string
	Darwin  []string
	Freebsd []string
}
type Package struct {
	Name         string
	Author       string
	Description  string
	Url          string
	Branch       string
	Install      []string
	Uninstall    []string
	Update       []string
	Config_paths []string
	Deps         *Deps
}

var environmentVariables = map[string]string{
	"PATH": home + ".local",
	"BIN":  home + ".local/bin",
}

func installPackages(pkgNames []string) {
	log(1, "Are you sure you would like to install the following packages:")
	for _, packageToInstall := range pkgNames {
		fmt.Println("        " + packageToInstall)
	}

	confirm("y", "(y/n)")

	for _, pkgName := range pkgNames {
		pkgInfoPath := installedPath + pkgName + ".json"
		url := "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json"

		if packageExists(pkgName) {
			log(4, "%s is already installed, can't install %s.", pkgName, pkgName)
			os.Exit(1)
		}

		log(1, "Getting package info for %s...", pkgName)
		log(1, "URL: %s", url)

		pkg, pkgFile := getPkgFromNet(pkgName)

		log(1, "Checking dependencies for %s...", pkgName)
		deps := getDeps(pkg)
		if deps != nil {
			log(1, "Dependencies: %s", strings.Join(deps, ", "))
			for _, dep := range deps {
				if checkIfCommandExists(dep) {
					log(0, "%s found!", dep)
				} else {
					log(4, "%s is either not installed or not in PATH. Please install it with your operating system's package manager.", dep)
					os.Exit(1)
				}
			}
		} else {
			log(1, "No dependencies found.")
		}

		initDirs("Making required directories for %s...", pkgName)

		log(1, "Writing package info for %s...", pkgName)
		newFile(pkgInfoPath, pkgFile, "An error occurred while writing package information for %s", pkgName)

		cloneRepo(pkg)

		log(1, "Running install commands for %s...", pkgName)
		runCommands(pkg.Install, pkg)

		log(0, "Installed %s successfully!\n", pkgName)
	}
}

func uninstallPackages(pkgNames []string) {
	log(1, "Are you sure you would like to uninstall the following packages:")
	for _, packageToUninstall := range pkgNames {
		fmt.Println("        " + packageToUninstall)
	}

	confirm("y", "(y/n)")

	for _, pkgName := range pkgNames {
		if !packageExists(pkgName) {
			log(3, "%s is not installed, so it can't be uninstalled", pkgName)
			continue
		}

		pkg := readAndLoad(pkgName)

		if purge {
			log(1, "Deleting configuration files for %s...", pkgName)
			for _, path := range pkg.Config_paths {
				log(1, "Deleting configuration path %s%s%s", textFx["BOLD"], path, RESETCOL)
				delPath(3, home+path, "An error occurred while deleting configuration files for %s", pkgName)
			}
		}

		log(1, "Running uninstall commands for %s...", pkgName)
		runCommands(pkg.Uninstall, pkg)

		log(1, "Deleting source files for %s...", pkgName)
		delPath(3, srcPath+pkgName, "An error occurred while deleting source files for %s", pkgName)

		log(1, "Deleting info file for %s...", pkgName)
		delPath(3, installedPath+pkgName+".json", "An error occurred while deleting info file for package %s", pkgName)

		log(0, "Successfully uninstalled %s.\n", pkgName)
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

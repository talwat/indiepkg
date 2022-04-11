package main

import (
	"fmt"
	"os"
	"strings"
)

var home string = os.Getenv("HOME") + "/"

var mainPath string = home + ".indiepkg/"
var srcPath string = mainPath + "data/package_src/"
var installedPath string = mainPath + "data/installed_packages/"
var bin string = home + ".local/bin/"

type Bin struct {
	Installed []string
	In_source []string
}
type Commands struct {
	Install   []string
	Uninstall []string
	Update    []string
}
type OSCommands struct {
	All    *Commands
	Linux  *Commands
	Darwin *Commands
}
type Deps struct {
	All    []string
	Linux  []string
	Darwin []string
}
type Package struct {
	Name         string
	Author       string
	Description  string
	Url          string
	Branch       string
	Bin          *Bin
	Deps         *Deps
	Commands     *OSCommands
	Config_paths []string
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
		pkgDisplayName := bolden(pkgName)

		pkgInfoPath := installedPath + pkgName + ".json"
		url := "https://raw.githubusercontent.com/talwat/indiepkg/main/packages/" + pkgName + ".json"

		if packageExists(pkgName) {
			log(4, "%s is already installed, can't install %s.", pkgDisplayName, pkgDisplayName)
			os.Exit(1)
		}

		log(1, "Getting package info for %s...", pkgDisplayName)
		log(1, "URL: %s", url)

		pkg, pkgFile := getPkgFromNet(pkgName)

		log(1, "Checking dependencies for %s...", pkgDisplayName)
		deps := getDeps(pkg)
		if deps != nil {
			log(1, "Dependencies: %s", strings.Join(deps, ", "))
			for _, dep := range deps {
				if checkIfCommandExists(dep) {
					log(0, "%s found!", bolden(dep))
				} else {
					log(4, "%s is either not installed or not in PATH. Please install it with your operating system's package manager.", bolden(dep))
					os.Exit(1)
				}
			}
		} else {
			log(1, "No dependencies found.")
		}

		initDirs("Making required directories for %s...", pkgDisplayName)

		log(1, "Writing package info for %s...", pkgDisplayName)
		newFile(pkgInfoPath, pkgFile, "An error occurred while writing package information for %s", pkgName)

		cloneRepo(pkg)

		if len(pkg.Bin.In_source) > 0 {
			log(1, "Copying binary files for %s...", pkgDisplayName)
			for i := range pkg.Bin.In_source {
				srcDir := srcPath + pkgName + "/" + pkg.Bin.In_source[i]
				destDir := bin + pkg.Bin.Installed[i]
				log(1, "Copying %s to %s...", bolden(srcDir), bolden(destDir))
				copyFile(srcDir, destDir)
				log(1, "Making %s executable...", bolden(destDir))
				changePerms(destDir, 0770)
			}
		}

		cmds := getInstCmd(pkg)

		if len(cmds) > 0 {
			log(1, "Running install commands for %s...", pkgDisplayName)
			runCommands(cmds, pkg, srcPath+pkg.Name)
		}

		log(0, "Installed %s successfully!\n", pkgDisplayName)
	}
}

func uninstallPackages(pkgNames []string) {
	log(1, "Are you sure you would like to uninstall the following packages:")
	for _, packageToUninstall := range pkgNames {
		fmt.Println("        " + packageToUninstall)
	}

	confirm("y", "(y/n)")

	for _, pkgName := range pkgNames {
		pkgDisplayName := bolden(pkgName)

		if !packageExists(pkgName) {
			log(3, "%s is not installed, so it can't be uninstalled", pkgDisplayName)
			continue
		}

		pkg := readAndLoad(pkgName)

		if purge {
			log(1, "Deleting configuration files for %s...", pkgDisplayName)
			for _, path := range pkg.Config_paths {
				log(1, "Deleting configuration path %s", bolden(home+path))
				delPath(3, home+path, "An error occurred while deleting configuration files for %s", pkgDisplayName)
			}
		}

		if len(pkg.Bin.Installed) > 0 {
			log(1, "Removing binary files for %s...", pkgDisplayName)
			for _, path := range pkg.Bin.Installed {
				log(1, "Removing %s", bolden(bin+path))
				delPath(4, bin+path, "An error occurred while removing binary files for %s", pkgDisplayName)
			}
		}

		cmds := getUninstCmd(pkg)

		if len(cmds) > 0 {
			log(1, "Running uninstall commands for %s...", pkgDisplayName)
			runCommands(cmds, pkg, srcPath+pkg.Name)
		}

		log(1, "Deleting source files for %s...", pkgDisplayName)
		delPath(3, srcPath+pkgName, "An error occurred while deleting source files for %s", pkgName)

		log(1, "Deleting info file for %s...", pkgDisplayName)
		delPath(3, installedPath+pkgName+".json", "An error occurred while deleting info file for package %s", pkgName)

		log(0, "Successfully uninstalled %s.\n", pkgDisplayName)
	}
}

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
var config string = mainPath + "config/"

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

func installPkgs(pkgNames []string) {
	initDirs(false)

	log(1, "Are you sure you would like to install the following packages:")
	for _, pkgToInstall := range pkgNames {
		fmt.Println("        " + pkgToInstall)
	}

	confirm("y", "(y/n)")

	for _, pkgName := range pkgNames {
		pkgDispName := bolden(pkgName)

		if pkgExists(pkgName) {
			log(4, "%s is already installed, can't install %s.", pkgDispName, pkgDispName)
			os.Exit(1)
		}

		pkg := downloadPkg(pkgName, true)

		log(1, "Checking dependencies for %s...", pkgDispName)
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

		cloneRepo(pkg)

		cmds := getInstCmd(pkg)

		runCmds(cmds, pkg, srcPath+pkg.Name, "install")

		if len(pkg.Bin.In_source) > 0 {
			log(1, "Copying binary files for %s...", pkgDispName)
			for i := range pkg.Bin.In_source {
				srcDir := srcPath + pkgName + "/" + pkg.Bin.In_source[i]
				destDir := bin + pkg.Bin.Installed[i]
				log(1, "Copying %s to %s...", bolden(srcDir), bolden(destDir))
				copyFile(srcDir, destDir)
				log(1, "Making %s executable...", bolden(destDir))
				changePerms(destDir, 0770)
			}
		}

		log(0, "Installed %s successfully!\n", pkgDispName)
	}
}

func uninstallPkgs(pkgNames []string) {
	log(1, "Are you sure you would like to uninstall the following packages:")
	for _, pkgToUninstall := range pkgNames {
		fmt.Println("        " + pkgToUninstall)
	}

	confirm("y", "(y/n)")

	for _, pkgName := range pkgNames {
		pkgDispName := bolden(pkgName)

		if !pkgExists(pkgName) {
			log(3, "%s is not installed, so it can't be uninstalled", pkgDispName)
			continue
		}

		pkg := readLoad(pkgName)

		if purge {
			log(1, "Deleting configuration files for %s...", pkgDispName)
			for _, path := range pkg.Config_paths {
				log(1, "Deleting configuration path %s", bolden(home+path))
				delPath(3, home+path, "An error occurred while deleting configuration files for %s", pkgDispName)
			}
		}

		if len(pkg.Bin.Installed) > 0 {
			log(1, "Removing binary files for %s...", pkgDispName)
			for _, path := range pkg.Bin.Installed {
				log(1, "Removing %s", bolden(bin+path))
				delPath(4, bin+path, "An error occurred while removing binary files for %s", pkgDispName)
			}
		}

		cmds := getUninstCmd(pkg)

		runCmds(cmds, pkg, srcPath+pkg.Name, "uninstall")

		log(1, "Deleting source files for %s...", pkgDispName)
		delPath(3, srcPath+pkgName, "An error occurred while deleting source files for %s", pkgName)

		log(1, "Deleting info file for %s...", pkgDispName)
		delPath(3, installedPath+pkgName+".json", "An error occurred while deleting info file for package %s", pkgName)

		log(0, "Successfully uninstalled %s.\n", pkgDispName)
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func installPkgs(pkgNames []string) {
	log(1, "Are you sure you would like to install the following packages:")
	for _, pkgToInstall := range pkgNames {
		fmt.Println("        " + pkgToInstall)
	}

	confirm("y", "(y/n)")

	initDirs(false)
	loadConfig()

	for _, pkgName := range pkgNames {
		pkgDispName := bolden(pkgName)

		if pkgExists(pkgName) {
			if force {
				log(3, "%s is already installed, but force is on, so continuing.", pkgDispName)
			} else {
				log(4, "%s is already installed, can't install %s.", pkgDispName, pkgDispName)
				os.Exit(1)
			}
		}

		pkg := downloadPkg(pkgName, true)

		if !noDeps {
			log(1, "Checking dependencies for %s...", pkgDispName)
			deps := getDeps(pkg)
			if deps != nil {
				log(1, "Dependencies: %s", strings.Join(deps, ", "))
				for _, dep := range deps {
					if checkIfCommandExists(dep) {
						log(0, "%s found!", bolden(dep))
					} else if force {
						log(3, "%s not found, but force is set, so continuing.", bolden(dep))
					} else {
						log(4, "%s is either not installed or not in PATH. Please install it with your operating system's package manager.", bolden(dep))
						os.Exit(1)
					}
				}
			} else {
				log(1, "No dependencies found.")
			}
		} else {
			log(3, "Skipping dependency check because nodeps is set to true.")
		}

		cloneRepo(pkg)

		cmds := getInstCmd(pkg)

		runCmds(cmds, pkg, srcPath+pkg.Name, "install")

		copyBins(pkg)

		log(0, "Installed %s successfully!\n", pkgDispName)
		getNotes(pkg)
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
			if force {
				log(3, "%s is not installed, but force is on, so continuing.", pkgDispName)
			} else {
				log(3, "%s is not installed, so it can't be uninstalled", pkgDispName)
				os.Exit(1)
			}
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
				log(1, "Removing %s", bolden(binPath+path))
				delPath(4, binPath+path, "An error occurred while removing binary files for %s", pkgDispName)
			}
		}

		cmds := getUninstCmd(pkg)

		runCmds(cmds, pkg, srcPath+pkg.Name, "uninstall")

		log(1, "Deleting source files for %s...", pkgDispName)
		delPath(3, srcPath+pkgName, "An error occurred while deleting source files for %s", pkgName)

		log(1, "Deleting info file for %s...", pkgDispName)
		delPath(3, infoPath+pkgName+".json", "An error occurred while deleting info file for package %s", pkgName)

		log(0, "Successfully uninstalled %s.\n", pkgDispName)
	}
}

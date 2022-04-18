package main

import (
	"os"
)

func installPkgs(pkgNames []string) {
	displayPkgs(pkgNames, "install")

	fullInit()

	for _, pkgName := range pkgNames {
		chapLog("=>", "VIOLET", "Installing %s", pkgName)
		chapLog("==>", "BLUE", "Getting package info")
		log(1, "Reading package info for %s...", bolden(pkgName))
		pkgFile := findPkg(pkgName)
		pkg := loadPkg(pkgFile, pkgName)
		cmds := getInstCmd(pkg)
		pkgDispName := bolden(pkgName)

		chapLog("==>", "BLUE", "Running checks")
		log(1, "Checking if %s is already installed...", pkgDispName)

		if pkgExists(pkgName) {
			if force {
				log(3, "%s is already installed, but force is on, so continuing.", pkgDispName)
			} else {
				log(4, "%s is already installed, can't install %s.", pkgDispName, pkgDispName)
				os.Exit(1)
			}
		}

		checkDeps(pkg, pkgName)

		if pkg.Download == nil {
			chapLog("==>", "BLUE", "Cloning source code")
			log(1, "Making sure %s is not already cloned...", pkgDispName)
			delPath(3, tmpSrcPath+pkg.Name, "An error occurred while deleting temporary source files for %s", pkgName)
			cloneRepo(pkg, tmpSrcPath)
		} else {
			doDirectDownload(pkg, pkgName, tmpSrcPath)
		}

		if len(cmds) > 0 {
			chapLog("==>", "BLUE", "Compiling")
			runCmds(cmds, pkg, tmpSrcPath+pkg.Name, "install")
		}

		chapLog("==>", "BLUE", "Installing")
		copyBins(pkg, tmpSrcPath)
		copyManpages(pkg, tmpSrcPath)
		mvPath(tmpSrcPath+pkg.Name, srcPath+pkg.Name)
		writeLoadPkg(pkg.Name, pkgFile, false)

		chapLog("=>", "GREEN", "Success")
		log(0, "Installed %s successfully!", pkgDispName)
		getNotes(pkg)
	}
}

func uninstallPkgs(pkgNames []string) {
	displayPkgs(pkgNames, "uninstall")

	fullInit()

	for _, pkgName := range pkgNames {
		chapLog("=>", "VIOLET", "Uninstalling %s", pkgName)
		pkgDispName := bolden(pkgName)

		chapLog("==>", "BLUE", "Running checks & getting info")
		if !pkgExists(pkgName) {
			if force {
				log(3, "%s is not installed, but force is on, so continuing.", pkgDispName)
			} else {
				log(3, "%s is not installed, so it can't be uninstalled", pkgDispName)
				os.Exit(1)
			}
		}

		pkg := readLoad(pkgName)

		chapLog("==>", "BLUE", "Deleting binary & configuration files")
		if purge {
			log(1, "Deleting configuration files for %s...", pkgDispName)
			for _, path := range pkg.Config_paths {
				log(1, "Deleting configuration path %s", bolden(home+path))
				delPath(3, home+path, "An error occurred while deleting configuration files for %s", pkgDispName)
			}
		}

		if len(pkg.Bin.Installed) > 0 {
			log(1, "Deleting binary files for %s...", pkgDispName)
			for _, path := range pkg.Bin.Installed {
				log(1, "Deleting %s", bolden(binPath+path))
				delPath(4, binPath+path, "An error occurred while deleting binary files for %s", pkgDispName)
			}
		}

		chapLog("==>", "BLUE", "Running uninstall commands")
		cmds := getUninstCmd(pkg)
		runCmds(cmds, pkg, srcPath+pkg.Name, "uninstall")

		chapLog("==>", "BLUE", "Deleting info & source")
		log(1, "Deleting source files for %s...", pkgDispName)
		delPath(3, srcPath+pkgName, "An error occurred while deleting source files for %s", pkgName)

		log(1, "Deleting info file for %s...", pkgDispName)
		delPath(3, infoPath+pkgName+".json", "An error occurred while deleting info file for package %s", pkgName)

		chapLog("=>", "GREEN", "Success")
		log(0, "Successfully uninstalled %s.", pkgDispName)
	}
}

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func pkgExists(pkgName string) bool {
	packageDisplayName := bolden(pkgName)

	infoInstalled := pathExists(infoPath+pkgName+".json", "package info for %s", packageDisplayName)
	srcInstalled := pathExists(srcPath+pkgName, "package source for %s", packageDisplayName)

	if infoInstalled && srcInstalled {
		return true
	} else if !infoInstalled && !srcInstalled {
		return false
	} else {
		log(4, "Package info or source for %s exists, but not both. Please run %sindiepkg sync%s.", packageDisplayName, textFx["BOLD"], RESETCOL)
		return false
	}
}

func runCmds(cmds []string, pkg Package, path string, cmdsLabel string) {
	debugLog("Work dir: %s", path)
	if len(cmds) > 0 {
		log(1, "Running %s commands for %s...", cmdsLabel, pkg.Name)
		for _, command := range cmds {
			logNoNewline(1, "Running command %s", bolden(command))
			runCommandRealTime(path, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
		}
	}
}

func initDirs(reset bool) {
	if reset {
		confirm("y", "Are you sure you want to reset the directories? This will reset your custom configuration & sources file. (y/n)")
	}

	log(1, "Making required directories & files...")
	newDir(srcPath, "An error occurred while creating sources directory")
	newDir(tmpSrcPath, "An error occurred while creating temporary sources directory")
	newDir(infoPath, "An error occurred while creating info directory")
	newDir(configPath, "An error occurred while creating config directory")

	if !pathExists(configPath+"config.json", "config file") || reset {
		log(1, "Creating config file...")
		newFile(configPath+"config.toml", defaultConf, "An error occurred while creating config file")
	}

	if !pathExists(configPath+"sources.txt", "sources file") || reset {
		log(1, "Creating sources file...")
		newFile(configPath+"sources.txt", defaultSources, "An error occurred while creating sources file")
	}
}

func getDeps(pkg Package) []string {
	if pkg.Deps != nil {
		fullDepsList := pkg.Deps.All
		switch runtime.GOOS {
		case "darwin":
			debugLog("Getting dependencies specifically for darwin...")
			fullDepsList = append(fullDepsList, pkg.Deps.Darwin...)
		case "linux":
			debugLog("Getting dependencies specifically for linux...")
			fullDepsList = append(fullDepsList, pkg.Deps.Linux...)
		default:
			log(3, "Unknown OS: %s", runtime.GOOS)
		}
		return fullDepsList
	}
	return nil
}

func checkDeps(pkg Package, pkgName string) {
	pkgDispName := bolden(pkgName)
	if !noDeps {
		log(1, "Getting dependencies for %s...", pkgDispName)
		deps := getDeps(pkg)

		log(1, "Checking dependencies for %s...", pkgDispName)
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
}

func parseSources() []string {
	log(1, "Reading sources file...")
	sourcesFile := readFile(configPath+"sources.txt", "An error occurred while reading sources file")

	if sourcesFile == defaultSources {
		debugLog("Default sources file detected.")
		return []string{"https://raw.githubusercontent.com/talwat/indiepkg/main/packages/"}
	}
	log(1, "Parsing sources file...")
	var finalList []string

	for _, line := range strings.Split(sourcesFile, "\n") {
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		finalList = append(finalList, line)
	}

	return finalList
}

func getNotes(pkg Package) {
	if len(pkg.Notes) > 0 {
		log(1, bolden("Important note!"))
		for _, note := range pkg.Notes {
			fmt.Println("        " + note)
		}
	}
}

func displayPkgs(pkgNames []string, action string) {
	log(1, "Are you sure you would like to %s the following packages:", bolden(action))
	for _, pkgToDisplay := range pkgNames {
		fmt.Println("        " + pkgToDisplay)
	}

	confirm("y", "(y/n)")
}

func fullInit() {
	chapLog("=>", "VIOLET", "Initializing")
	initDirs(false)
	loadConfig()
}

package main

import (
	"fmt"
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

func copyBins(pkg Package, srcPath string) {
	pkgDispName := bolden(pkg.Name)
	log(1, "Making binary directory...")
	newDir(binPath, "An error occurred while creating binaries directory")
	if len(pkg.Bin.In_source) > 0 {
		log(1, "Copying files for %s...", pkgDispName)
		for i := range pkg.Bin.In_source {
			srcDir := srcPath + pkg.Name + "/" + pkg.Bin.In_source[i]
			destDir := binPath + pkg.Bin.Installed[i]
			log(1, "Copying %s to %s...", bolden(srcDir), bolden(destDir))
			copyFile(srcDir, destDir)
			log(1, "Making %s executable...", bolden(destDir))
			changePerms(destDir, 0770)
		}
	}
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

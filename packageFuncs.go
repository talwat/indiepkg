package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func loadPkg(packageFile string, pkgName string) Package {
	var pkg Package

	debugLog("Finding environment variables...")
	keySlice := make([]string, 0)
	for key := range environmentVariables {
		keySlice = append(keySlice, key)
	}

	debugLog("Replacing environment variables...")
	for _, key := range keySlice {
		packageFile = strings.Replace(packageFile, ":("+key+"):", environmentVariables[key], -1)
	}
	err := json.Unmarshal([]byte(packageFile), &pkg)
	errorLog(err, 4, "An error occurred while loading package info for %s", pkgName)
	return pkg
}

func readLoad(pkgName string) Package {
	packageDisplayName := bolden(pkgName)

	log(1, "Reading package info for %s...", packageDisplayName)
	pkgFile := readFile(installedPath+pkgName+".json", "An error occurred while reading package %s", packageDisplayName)

	log(1, "Loading package info for %s...", packageDisplayName)
	pkg := loadPkg(pkgFile, fmt.Sprintf("An error occurred while loading package information for %s", packageDisplayName))

	return pkg
}

func pkgExists(pkgName string) bool {
	packageDisplayName := bolden(pkgName)

	infoInstalled := pathExists(installedPath+pkgName+".json", "package info for %s", packageDisplayName)
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
	newDir(installedPath, "An error occurred while creating info directory")
	newDir(configPath, "An error occurred while creating config directory")

	if !pathExists(configPath+"config.json", "config file") || reset {
		log(1, "Creating config file...")
		newFile(configPath+"config.json", defaultConf, "An error occurred while creating config file")
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

func cloneRepo(pkg Package) {
	log(1, "Cloning source code for %s...", bolden(pkg.Name))
	if pkg.Branch == "" {
		debugLog("Cloning to %s", bolden(srcPath+pkg.Name))

		_, err := git.PlainClone(srcPath+pkg.Name, false, &git.CloneOptions{
			URL:      pkg.Url,
			Progress: os.Stdout,
		})

		errorLog(err, 4, "An error occurred while cloning repository for %s", bolden(pkg.Name))
	} else {
		log(1, "Getting branch %s%s%s...", textFx["BOLD"], pkg.Branch, RESETCOL)
		debugLog("Cloning to %s on branch %s", srcPath+pkg.Name, pkg.Branch)
		_, err := git.PlainClone(srcPath+pkg.Name, false, &git.CloneOptions{
			URL:           pkg.Url,
			Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", pkg.Branch)),
			SingleBranch:  true,
		})

		errorLog(err, 4, "An error occurred while cloning repository for %s", bolden(pkg.Name))
	}
}

func pullRepo(pkgName string) error {
	var err error
	r, err := git.PlainOpen(srcPath + pkgName)
	errorLog(err, 4, "An error occurred while opening repository for %s", bolden(pkgName))

	w, err := r.Worktree()
	errorLog(err, 4, "An error occurred while getting worktree for %s", bolden(pkgName))

	debugLog("Pulling %s", bolden(srcPath+pkgName))
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})

	if err.Error() == "already up-to-date" {
		log(0, "%s already up to date.", bolden(pkgName))
	} else {
		errorLog(err, 4, "An error occurred while pulling repository for %s", bolden(pkgName))
	}
	return err
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

func copyBins(pkg Package) {
	pkgDispName := bolden(pkg.Name)
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

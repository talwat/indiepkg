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

func loadPackage(packageFile string, pkgName string) Package {
	var pkg Package

	log(1, "Finding environment variables...")
	keySlice := make([]string, 0)
	for key := range environmentVariables {
		keySlice = append(keySlice, key)
	}

	log(1, "Replacing environment variables...")
	for _, key := range keySlice {
		packageFile = strings.Replace(packageFile, ":("+key+"):", environmentVariables[key], -1)
	}
	err := json.Unmarshal([]byte(packageFile), &pkg)
	errorLog(err, 4, "An error occurred while loading package info for %s", pkgName)
	return pkg
}

func readAndLoad(packageName string) Package {
	log(1, "Reading package info for %s...", packageName)
	pkgFile := readFile(installedPath+packageName+".json", "An error occurred while reading package %s", packageName)

	log(1, "Loading package info for %s...", packageName)
	pkg := loadPackage(pkgFile, fmt.Sprintf("An error occurred while loading package information for %s", packageName))

	return pkg
}

func packageExists(pkgName string) bool {
	infoInstalled := pathExists(installedPath+pkgName+".json", "An error occurred while checking if package info for %s exists", pkgName)

	srcInstalled := pathExists(srcPath+pkgName, "An error occurred while checking if package source for %s exists", pkgName)

	if infoInstalled && srcInstalled {
		return true
	} else if !infoInstalled && !srcInstalled {
		return false
	} else {
		log(4, "Package info or source for %s exists, but not both. Please run %sindiepkg repair%s", pkgName, textFx["BOLD"], RESETCOL)
		return false
	}
}

func runCommands(commands []string, pkg Package, path string) {
	for _, command := range commands {
		log(1, "Running command %s%s%s...", textFx["BOLD"], command, RESETCOL)
		runCommand(path, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
	}
}

func initDirs(msg string, params ...interface{}) {
	log(1, fmt.Sprintf(msg, params...))
	newDirSilent(srcPath)
	newDirSilent(installedPath)
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

func getPkgFromNet(pkgName string) (Package, string) {
	packageFile, err := viewFile("https://raw.githubusercontent.com/talwat/indiepkg/main/packages/"+pkgName+".json", "An error occurred while getting package information for %s", pkgName)

	if errIs404(err) {
		log(4, "Package %s not found.", bolden(pkgName))
		os.Exit(1)
	}

	errorLog(err, 4, "An error occurred while getting package information for %s", bolden(pkgName))

	pkg := loadPackage(packageFile, pkgName)

	return pkg, packageFile
}

func errIs404(err error) bool {
	return err != nil && strings.Contains(err.Error(), "404")
}

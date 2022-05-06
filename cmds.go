package main

import (
	"os"
	"runtime"
	"strings"
)

func listPkgs() {
	installedPkgs := make([]string, 0)
	files := dirContents(infoPath, "An error occurred while getting list of installed packages")

	if len(files) == 0 {
		log(1, "No packages installed.")
		os.Exit(0)
	}

	for _, file := range files {
		installedPkgs = append(installedPkgs, strings.ReplaceAll(file.Name(), ".json", ""))
	}

	rawLogf(strings.Join(installedPkgs, "\n") + "\n")
}

func sync() {
	fullInit()

	chapLog("==>", "", "Getting packages to sync")

	log(1, "Getting list of package files...")

	dirs := dirContents(srcPath, "An error occurred while getting list of source files")

	log(1, "Getting missing package info...")

	var pkgInfoToSync []string

	for _, dir := range dirs {
		pkgName := strings.ReplaceAll(dir.Name(), ".json", "")
		infoExists := pathExists(infoPath+pkgName+".json", "An error occurred while checking if %s is properly installed", pkgName)

		if !infoExists && dir.IsDir() {
			pkgInfoToSync = append(pkgInfoToSync, pkgName)
		}
	}

	log(1, "Getting list of source directories...")

	infoFiles := dirContents(infoPath, "An error occurred while getting list of info files")

	log(1, "Getting missing source directories...")

	var pkgSrcToSync []string

	for _, infoFile := range infoFiles {
		pkgName := strings.ReplaceAll(infoFile.Name(), ".json", "")
		srcExists := pathExists(srcPath+pkgName, "An error occurred while checking if %s is properly installed", pkgName)

		if !srcExists && !infoFile.IsDir() {
			pkgSrcToSync = append(pkgSrcToSync, pkgName)
		}
	}

	chapLog("=>", "", "Syncing packages...")

	for _, pkgToSync := range pkgInfoToSync {
		chapLog("==>", "", "Downloading info for %s", pkgToSync)
		downloadPkg(pkgToSync)
	}

	for _, pkgToSync := range pkgSrcToSync {
		pkg := readLoad(pkgToSync)
		if pkg.Download == nil {
			chapLog("==>", "", "Cloning source for %s", pkgToSync)
			clonePkgRepo(readLoad(pkgToSync), srcPath)
		} else {
			chapLog("==>", "", "Downloading file for %s", pkgToSync)
			doDirectDownload(pkg, pkgToSync, srcPath)
		}
	}

	if len(pkgInfoToSync) > 0 || len(pkgSrcToSync) > 0 {
		chapLog("=>", "GREEN", "Success")
		log(0, "Successfully synced info for %s and source for %s!", strings.Join(pkgInfoToSync, ", "), strings.Join(pkgSrcToSync, ", "))
	} else {
		chapLog("=>", "CYAN", "Nothing synced")
		log(1, "Nothing synced.")
	}
}

func infoPkg(pkgName string) {
	pkg := loadPkg(getPkgInfo(pkgName, isURL(pkgName)), pkgName)

	safePrintVal := func(name string, val string) {
		if val != "" {
			log(1, "%s: %s", name, val)
		} else {
			log(1, "%s: Undefined", name)
		}
	}

	rawLogf("\n")
	log(1, "Name: %s", pkg.Name)
	log(1, "Author: %s", pkg.Author)
	log(1, "Description: %s", pkg.Description)
	safePrintVal("License", pkg.License)
	safePrintVal("Programming Language", pkg.Language)
	log(1, "Git URL: %s", pkg.URL)

	if deps := getDeps(pkg.Deps); deps != nil {
		log(1, "Dependencies: %s", strings.Join(deps, ", "))
	}

	if deps := getDeps(pkg.FileDeps); deps != nil {
		log(1, "File dependencies: %s", strings.Join(deps, ", "))
	}

	getNotes(pkg)
}

func rmData(pkgNames []string) {
	log(3, "Warning: This will remove the data for the selected packages stored in %s", mainPath)
	log(3, "This will %snot%s run the uninstall commands.", textFx["BOLD"], RESETCOL)
	log(3, "You should only use this in case a package installation has failed at a certain step, or you want to separate an installed package from indiepkg.")
	displayPkgs(pkgNames, "remove the data for")

	for _, pkgName := range pkgNames {
		chapLog("=>", "", "Removing data for %s", pkgName)
		pkgDisplayName := bolden(pkgName)

		log(1, "Deleting source files for %s...", pkgDisplayName)
		delPath(false, srcPath+pkgName, "An error occurred while deleting source files for %s", pkgDisplayName)

		log(1, "Deleting info file for %s...", pkgDisplayName)
		delPath(false, infoPath+pkgName+".json", "An error occurred while deleting info file for %s", pkgDisplayName)

		log(0, "Successfully deleted the data for %s.\n", pkgDisplayName)
	}

	chapLog("=>", "GREEN", "Success")
	log(0, "Successfully deleted data.")
}

func search(query string) {
	initDirs(false)
	loadConfig()

	pkgs, _ := getPkgFromGh(query)

	rawLogf("\n")
	log(1, "Found %d packages:", len(pkgs))

	for _, pkg := range pkgs {
		rawLogf("        " + pkg.Name + " - " + pkg.Repo + "\n")
	}
}

func listAll() {
	initDirs(false)
	loadConfig()

	pkgs, _ := getAllPkgsFromGh()

	rawLogf("\n")
	log(1, "Found %d packages:", len(pkgs))

	for _, pkg := range pkgs {
		rawLogf("        " + pkg.Name + " - " + repoLabel(pkg.Repo, true) + "\n")
	}
}

func reClone() {
	loadConfig()
	log(1, "Resetting IndiePKG source directory...")
	delPath(true, indiePkgSrcDir, "An error occurred while deleting the IndiePKG source directory")

	cloneSrcRepo()
	log(0, "Successfully re-cloned IndiePKG source.")
}

func fetch() {
	log(1, "OS: %s", runtime.GOOS)
	log(1, "Arch: %s", runtime.GOARCH)
	log(1, "Go Version: %s", strings.TrimPrefix(runtime.Version(), "go"))
	log(1, "IndiePKG Version: %s", version)

	macOSVer, err := runCommand(".", "sw_vers", "-productVersion")
	if err == nil {
		log(1, "macOS version: %s", macOSVer)
	}

	bashVer, err := runCommand(".", "bash", "--version")
	if err == nil {
		log(1, "Bash version: %s", strings.Split(strings.TrimPrefix(strings.Split(bashVer, "\n")[0], "GNU bash, version "), "(")[0])
	} else {
		log(3, "Could not get bash version. Error: %s", err.Error())
	}

	if pathExists("/etc/os-release", "An error occurred while checking if /etc/os-release exists") {
		log(1, "OS-Release:")
		indent(readFile("/etc/os-release", "An error occurred while reading /etc/os-release"))
	}
}

func help2man() {
	log(1, "Compiling IndiePKG...")

	runCommand(".", "make")

	log(1, "Running help2man...")

	output, _ := runCommand(".", "help2man", "-h", "help", "-v", "raw-version", "./indiepkg")

	log(1, "Parsing generated manpage...")

	output = strings.ReplaceAll(output, ".TP", ".SS")
	output = strings.ReplaceAll(output, ".SS \"Commands:\"\n", "")
	output = strings.ReplaceAll(output, ".SS \"Developer & Debugging Commands:\"\n", "")

	log(1, "Writing final manpage...")

	newFile("indiepkg.1", output, "An error occurred while writing generated manpage")
}

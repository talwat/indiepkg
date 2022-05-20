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

	rawLog(strings.Join(installedPkgs, "\n") + "\n")
}

func infoPkg(pkgName string) {
	pkg := loadPkg(getPkgInfo(pkgName, isValidURL(pkgName)), pkgName)

	safePrintVal := func(name string, val string) {
		if val != "" {
			log(1, "%s: %s", name, val)
		} else {
			log(1, "%s: Undefined", name)
		}
	}

	rawLog("\n")
	log(1, "Name: %s", pkg.Name)
	log(1, "Author: %s", pkg.Author)
	log(1, "Description: %s", pkg.Description)
	safePrintVal("License", pkg.License)
	safePrintVal("Programming Language", pkg.Language)
	safePrintVal("Git URL", pkg.URL)

	if pkg.Bin != nil {
		log(1, "Binaries: %s", strings.Join(pkg.Bin.Installed, ", "))
	}

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
	log(3, "This will %snot%s run the uninstall commands.", textFx.Bold, RESETCOL)
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

	chapLog("=>", textCol.Green, "Success")
	log(0, "Successfully deleted data.")
}

func search(query string) {
	initDirs(false)
	loadConfig()

	pkgs, _ := getPkgFromGh(query)

	rawLog("\n")
	log(1, "Found %d packages:", len(pkgs))

	for _, pkg := range pkgs {
		rawLog("        " + pkg.Name + " - " + pkg.Repo + "\n")
	}
}

func listAll() {
	initDirs(false)
	loadConfig()

	pkgs, _ := getAllPkgsFromGh()

	rawLog("\n")
	log(1, "Found %d packages:", len(pkgs))

	for _, pkg := range pkgs {
		rawLog("        " + pkg.Name + " - " + repoLabel(pkg.Repo, true) + "\n")
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
	configExists := pathExists(configPath+"config.toml", "config file")

	if configExists {
		loadConfig()
	} else {
		log(3, "No config file found, skipping configuration related info.")
	}

	chapLog("=>", "", "OS & Environment")

	log(1, "OS: %s", runtime.GOOS)
	log(1, "Arch: %s", runtime.GOARCH)
	log(1, "Go Version: %s", strings.TrimPrefix(runtime.Version(), "go"))

	if isRoot() {
		log(3, "Running as %s", textCol.Red+"root"+RESETCOL)
	} else {
		log(1, "Running as %s", textCol.Green+"normal user"+RESETCOL)
	}

	macOSVer, err := runCommand(".", "sw_vers", "-productVersion")
	if err == nil {
		log(1, "macOS version: %s", macOSVer)
	}

	if pathExists("/etc/os-release", "os release") {
		log(1, "OS-Release:")
		indent(readFile("/etc/os-release", "An error occurred while reading /etc/os-release"))
	}

	kern, err := runCommand(".", "uname", "-rsp")
	if err != nil {
		log(3, "Could not get kernel info. Error: %s", err.Error())
	} else {
		log(1, "Kernel: %s", kern)
	}

	chapLog("=>", "", "IndiePKG Information")
	log(1, "IndiePKG Version: %s", version)

	if configExists {
		log(1, "IndiePKG Branch: %s", config.Updating.Branch)
	}

	chapLog("=>", "", "CLI Information")
	log(1, "Shell: %s", os.Getenv("SHELL"))
	log(1, "TERM: %s", os.Getenv("TERM"))

	bashVer, err := runCommand(".", "bash", "--version")
	if err == nil {
		log(1, "Bash version: %s", strings.Split(strings.TrimPrefix(strings.Split(bashVer, "\n")[0], "GNU bash, version "), "(")[0])
	} else {
		log(3, "Could not get bash version. Error: %s", err.Error())
	}

	zshVer, err := runCommand(".", "zsh", "--version")
	if err == nil {
		log(1, "Zsh version: %s", strings.Split(strings.TrimPrefix(zshVer, "zsh "), "(")[0])
	} else {
		log(3, "Could not get zsh version. Error: %s", err.Error())
	}
}

func help2man() {
	log(1, "Compiling IndiePKG...")

	_, err := runCommand(".", "make")

	errorLog(err, "An error occurred while compiling IndiePKG")

	log(1, "Running help2man...")

	output, _ := runCommand(".", "help2man", "-h", "help", "-v", "raw-version", "./indiepkg")

	log(1, "Parsing generated manpage...")

	output = strings.ReplaceAll(output, ".TP", ".SS")
	output = strings.ReplaceAll(output, ".SS \"Commands:\"\n", "")
	output = strings.ReplaceAll(output, ".SS \"Developer & Debugging Commands:\"\n", "")

	log(1, "Writing final manpage...")

	newFile("indiepkg.1", output, "An error occurred while writing generated manpage")
}

func reCompile(pkgNames []string) {
	displayPkgs(pkgNames, "re-compile")

	fullInit()

	for _, pkgName := range pkgNames {
		pkgDispName := bolden(pkgName)
		chapLog("==>", "", "Preparing for re-compilation of %s", pkgDispName)

		chapLog("===>", "", "Checking if already installed")
		log(1, "Checking if %s is already installed...", pkgDispName)

		if !pkgExists(pkgName) {
			if force {
				log(3, "%s is not installed, but force is on, so continuing.", pkgDispName)
			} else {
				errorLogRaw("%s is not installed, so it can't be re-compiled", pkgDispName)
			}
		}

		chapLog("===>", "", "Getting info")

		pkg := readLoad(pkgName)
		cmds := getInstCmd(pkg)

		chapLog("===>", "", "Checking dependencies")
		checkDeps(pkg)
		checkFileDeps(pkg)

		chapLog("==>", "", "Re-installing")

		if len(cmds) > 0 {
			chapLog("===>", "", "Compiling")
			runCmds(cmds, pkg, srcPath+pkg.Name, "install")
		}

		chapLog("===>", "", "Moving files")
		copyBins(pkg, srcPath)
		copyManpages(pkg, srcPath)

		chapLog("===>", textCol.Green, "Success")
		log(0, "Successfully re-compiled %s.", pkgDispName)
	}

	chapLog("=>", textCol.Green, "Success")
	log(0, "Successfully re-compiled all selected packages.")
}

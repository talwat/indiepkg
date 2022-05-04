package main

import (
	"os"
	"runtime"
	"strings"
)

func pkgExists(pkgName string) bool {
	packageDisplayName := bolden(pkgName)

	infoInstalled := pathExists(infoPath+pkgName+".json", "package info for %s", packageDisplayName)
	srcInstalled := pathExists(srcPath+pkgName, "package source for %s", packageDisplayName)

	switch {
	case infoInstalled && srcInstalled:
		return true
	case !infoInstalled && !srcInstalled:
		return false
	default:
		log(3, "Package info or source for %s exists, but not both. Please run %sindiepkg sync%s.", packageDisplayName, textFx["BOLD"], RESETCOL)

		return false
	}
}

func runCmds(cmds []string, pkg Package, path string, cmdsLabel string) {
	debugLog("Work dir: %s", path)

	if len(cmds) > 0 {
		log(1, "Running %s commands for %s...", cmdsLabel, pkg.Name)

		for _, command := range cmds {
			logNoNewline(1, "Running command %s", bolden(command))
			runCommandDot(path, strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
			rawLogf("\n")
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
	newDir(config.Paths.Prefix, "An error occurred while creating prefix directory")
	newDir(config.Paths.Prefix+"bin", "An error occurred while creating binary directory")
	newDir(config.Paths.Prefix+"share/man", "An error occurred while creating manpage directory")

	if !pathExists(configPath+"config.toml", "config file") || reset {
		log(1, "Creating config file...")
		newFile(configPath+"config.toml", defaultConf, "An error occurred while creating config file")
	}

	if !pathExists(configPath+"sources.txt", "sources file") || reset {
		log(1, "Creating sources file...")
		newFile(configPath+"sources.txt", defaultSources, "An error occurred while creating sources file")
	}

	if !pathExists(indiePkgSrcDir, "IndiePKG source directory") || reset {
		if reset {
			log(1, "Resetting IndiePKG source directory...")
			delPath(true, indiePkgSrcDir, "An error occurred while deleting the IndiePKG source directory")
		}

		cloneSrcRepo()
	}
}

func getDeps(deps *Deps) []string {
	if deps != nil {
		fullDepsList := deps.All

		switch runtime.GOOS {
		case "darwin":
			debugLog("Getting dependencies specifically for darwin...")

			fullDepsList = append(fullDepsList, deps.Darwin...)
		case "linux":
			debugLog("Getting dependencies specifically for linux...")

			fullDepsList = append(fullDepsList, deps.Linux...)
		default:
			log(3, "Unknown OS: %s", runtime.GOOS)
		}

		return fullDepsList
	}

	return nil
}

func getPkgNameFromURL(url string) string {
	log(1, "Getting package name from URL...")

	return strings.TrimSuffix(strings.Split(url, "/")[len(strings.Split(url, "/"))-1], ".json")
}

func checkDeps(pkg Package) {
	if pkgDispName := bolden(pkg.Name); !noDeps {
		log(1, "Getting dependencies for %s...", pkgDispName)

		deps := getDeps(pkg.Deps)

		log(1, "Checking dependencies for %s...", pkgDispName)

		if deps != nil {
			log(1, "Dependencies: %s", strings.Join(deps, ", "))

			for _, dep := range deps {
				switch {
				case checkIfCommandExists(dep):
					log(0, "%s found!", bolden(dep))
				case force:
					log(3, "%s not found, but force is set, so continuing.", bolden(dep))
				default:
					errorLogRaw("%s is either not installed or not in PATH. Please install it with your operating system's package manager", bolden(dep))
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

func checkFileDeps(pkg Package) {
	if pkgDispName := bolden(pkg.Name); !noDeps {
		log(1, "Getting file dependencies for %s...", pkgDispName)

		deps := getDeps(pkg.FileDeps)

		log(1, "Checking file dependencies for %s...", pkgDispName)

		if deps != nil {
			log(1, "File dependencies: %s", strings.Join(deps, ", "))

			for _, dep := range deps {
				switch {
				case pathExists(dep, bolden(dep)):
					log(0, "%s found!", bolden(dep))
				case force:
					log(3, "%s does not exist, but force is set, so continuing.", bolden(dep))
				default:
					errorLogRaw("%s does not exist, please install the package that provides it with your operating system's package manager.", bolden(dep))
					os.Exit(1)
				}
			}
		} else {
			log(1, "No file dependencies found.")
		}
	} else {
		log(3, "Skipping file dependency check because nodeps is set to true.")
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

	finalList := make([]string, 0)

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
		indent(strings.Join(pkg.Notes, "\n"))
	}
}

func displayPkgs(pkgNames []string, action string) {
	log(1, "Are you sure you would like to %s the following packages:", bolden(action))

	for _, pkgToDisplay := range pkgNames {
		indent(pkgToDisplay)
	}

	confirm("y", "(y/n)")
}

func checkRoot() {
	if isRoot() {
		if force {
			log(3, "Running as root, but force is set, so continuing.")

			return
		}

		errorLogRaw("You are running as root. This is dangerous, and is absolutely not recommended. Please try again as a normal user")
	}
}

func fullInit() {
	chapLog("=>", "", "Initializing")
	checkRoot()
	initDirs(false)
	loadConfig()
	autoUpdate()
}

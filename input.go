package main

import (
	"os"
	"strings"
)

const version = "0.31-alpha"

var purge, debug, assumeYes, force, noDeps bool = false, false, false, false, false

var optionToOthers, optionToOther bool = false, false

func checkFlag(flag string) {
	switch flag {
	case "-p", "--purge":
		purge = true
	case "-d", "--debug":
		debug = true
	case "-y", "--assumeyes":
		assumeYes = true
	case "-f", "--force":
		force = true
	case "-n", "--nodeps":
		noDeps = true
	default:
		errorLogRaw("Flag %s not found", bolden(flag))
		os.Exit(1)
	}
}

func checkCommand(other string, others []string, index int, args []string) {
	checkForOptions := func(errSpecify string, commandPartsCount int) {
		if len(others[index+commandPartsCount:]) < 1 {
			errorLogRaw("No %s specified", errSpecify)
			os.Exit(1)
		}
	}

	switch other {
	case "install":
		checkForOptions("package names", 1)

		optionToOthers = true

		installPkgs(others[index+1:])

	case "uninstall":
		checkForOptions("package names", 1)

		optionToOthers = true

		uninstallPkgs(others[index+1:])

	case "remove-data":
		checkForOptions("package names", 1)

		optionToOthers = true

		rmData(others[index+1:])

	case "upgrade":
		if len(others) <= index+1 {
			upgradeAllPackages()
		} else {
			optionToOthers = true

			upgradePackage(others[index+1:])
		}

	case "update":
		if len(args) <= index+1 {
			updateAllPackages()
		} else {
			optionToOthers = true

			updatePackage(others[index+1:])
		}

	case "info":
		checkForOptions("package name", 1)

		optionToOther = true

		infoPkg(others[index+1])

	case "sync":
		sync()

	case "list-all":
		listAll()

	case "re-clone":
		reClone()

	case "version":
		log(1, "Indiepkg Version %s", bolden(version))

	case "raw-version":
		rawLog(version)

	case "help":
		rawLog(helpMsg)

	case "list":
		listPkgs()

	case "init":
		initDirs(true)

	case "repo":
		checkForOptions("sub-command", 1)

		switch others[index+1] {
		case "add":
			checkForOptions("url", 2)
			addRepo(others[index+2])
		case "remove":
			checkForOptions("url", 2)
			rmRepo(others[index+2])
		case "list":
			listRepos()
		default:
			errorLogRaw("Sub-command %s not found", bolden(others[index+1]))
			os.Exit(1)
		}

	case "search":
		checkForOptions("query", 1)

		optionToOther = true

		search(others[index+1])

	case "indiepkg-update":
		updateIndiePKG()

	case "setup":
		setup()

	case "github-gen":
		optionToOthers = true

		checkForOptions("author", 1)
		checkForOptions("repo", 2)
		getRepoInfo(others[index+1], others[index+2])

	default:
		errorLogRaw("Command %s not found", bolden(other))
		os.Exit(1)
	}
}

func parseInput() {
	args := os.Args[1:]

	var (
		flags  []string
		others []string
	)

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		} else {
			others = append(others, arg)
		}
	}

	for _, flag := range flags {
		checkFlag(flag)
	}

	for i, other := range others {
		if !optionToOthers && !optionToOther {
			checkCommand(other, others, i, args)
		}

		optionToOther = true
	}

	if len(others) < 1 {
		log(1, "Indiepkg Version %s, run %sindiepkg help%s for usage.", bolden(version), textFx["BOLD"], RESETCOL)
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

const version = "0.22.3"

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
		log(4, "Flag %s not found.", bolden(flag))
		os.Exit(1)
	}
}

func checkCommand(other string, others []string, i int, args []string) {
	checkForOptions := func(errSpecify string, commandPartsCount int) {
		if len(others[i+commandPartsCount:]) < 1 {
			log(4, "No %s specified.", errSpecify)
			os.Exit(1)
		}
	}

	switch other {
	case "install":
		checkForOptions("package names", 1)
		optionToOthers = true
		installPkgs(others[i+1:])

	case "uninstall":
		checkForOptions("package names", 1)
		optionToOthers = true
		uninstallPkgs(others[i+1:])

	case "remove-data":
		checkForOptions("package names", 1)
		optionToOthers = true
		rmData(others[i+1:])

	case "upgrade":
		if len(others) <= i+1 {
			upgradeAllPackages()
		} else {
			optionToOthers = true
			upgradePackage(others[i+1:])
		}

	case "update":
		if len(args) <= i+1 {
			updateAllPackages()
		} else {
			optionToOthers = true
			updatePackage(others[i+1:])
		}

	case "info":
		checkForOptions("package name", 1)
		optionToOther = true
		infoPkg(others[i+1])

	case "sync":
		sync()

	case "version":
		log(1, "Indiepkg Version %s", bolden(version))

	case "help":
		fmt.Print(helpMsg)

	case "list":
		listPkgs()

	case "init":
		initDirs(true)

	case "repo":
		checkForOptions("sub-command", 1)
		switch others[i+1] {
		case "add":
			checkForOptions("url", 2)
			addRepo(others[i+2])
		case "remove":
			checkForOptions("url", 2)
			rmRepo(others[i+2])
		case "list":
			listRepos()
		default:
			log(4, "Sub-command %s not found.", bolden(others[i+1]))
			os.Exit(1)
		}

	case "search":
		checkForOptions("query", 1)
		optionToOther = true
		search(others[i+1])

	case "indiepkg-update":
		updateIndiePKG()

	default:
		log(4, "Command %s not found.", bolden(other))
		os.Exit(1)
	}
}

func parseInput() {
	args := os.Args[1:]
	var flags []string
	var others []string

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

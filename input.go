package main

import (
	"fmt"
	"os"
	"strings"
)

const version = "0.10-beta"

var purge, debug, assumeYes bool = false, false, false

var optionToOthers, optionToOther bool = false, false

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
		switch flag {
		case "-p", "--purge":
			purge = true
		case "-d", "--debug":
			debug = true
		case "-y", "--assumeYes":
			assumeYes = true
		default:
			log(1, "Flag %s not found.", bolden(flag))
		}
	}

	for i, other := range others {
		if !optionToOthers && !optionToOther {
			switch other {
			case "install":
				optionToOthers = true
				installPkgs(others[i+1:])

			case "uninstall":
				optionToOthers = true
				uninstallPkgs(others[i+1:])

			case "remove-data":
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
				optionToOther = true
				infoPkg(others[i+1])

			case "sync":
				sync()

			case "version":
				log(1, "Indiepkg Version %s", bolden(version))

			case "help":
				fmt.Printf(helpMsg)

			case "list":
				listPkgs()

			default:
				log(1, "Command %s not found.", bolden(other))
			}
		}

		optionToOther = true
	}

	if len(others) < 1 {
		log(1, "Indiepkg Version %s, run %sindiepkg help%s for usage.", bolden(version), textFx["BOLD"], RESETCOL)
	}
}

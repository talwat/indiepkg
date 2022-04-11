package main

import (
	"fmt"
	"os"
	"strings"
)

const version = "0.6"

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
			log(1, "Flag %s not found.", flag)
		}
	}

	for i, other := range others {
		if !optionToOthers && !optionToOther {
			switch other {
			case "install":
				optionToOthers = true
				installPackages(others[i+1:])

			case "uninstall":
				optionToOthers = true
				uninstallPackages(others[i+1:])

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
				infoPackage(others[i+1])

			case "repair":
				repair()

			case "version":
				log(1, "Indiepkg Version %s", version)

			case "help":
				fmt.Printf(`Usage: indiepkg [<option>...] <command>

Commands:
  install <packages>...
  uninstall <packages>...
  update <packages>...
  upgrade <packages>...
  info <package>
  repair
  version

Options:
  -p, --purge
  -d, --debug
  -y, --assumeYes

Examples:
  indiepkg install my-pkg
  indiepkg uninstall other-pkg
  indiepkg upgrade third-pkg
`)

			case "list":
				listPackages()

			default:
				log(1, "Command %s not found.", other)
			}
		}

		optionToOther = true
	}

	if len(others) < 1 {
		log(1, "Indiepkg Version %s, run %sindiepkg help%s for usage.", version, textFx["BOLD"], RESETCOL)
	}
}

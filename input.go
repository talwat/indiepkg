package main

import (
	"os"
	"strings"
)

var purge bool = false
var debug bool = false

func parseInput() {
	args := os.Args[1:]
	var commandSelected bool = false

	for i, arg := range args {
		if strings.HasPrefix(arg, "-") { // Flags
			switch arg {
			case "-p", "--purge":
				purge = true
			case "-d", "--debug":
				debug = true
			default:
				log(1, "Flag %s not found.", arg)
			}
		} else if !commandSelected { // Commands
			switch arg {
			case "install":
				installPackages(args[i+1:])
			case "uninstall":
				uninstallPackages(args[i+1:])
			case "upgrade":
				if len(args) <= i+1 {
					upgradeAllPackages()
				} else {
					upgradePackage(args[i+1:])
				}
			case "update":
				if len(args) <= i+1 {
					updateAllPackages()
				} else {
					updatePackage(args[i+1:])
				}
			case "info":
				infoPackage(args[i+1])
			case "repair":
				repair()
			case "version":
				log(1, "Indiepkg Version 0.1.3")
			case "help":
				log(1, "Help menu not done yet.")
			case "list":
				listPackages()
			default:
				log(1, "Command %s not found.", arg)
			}
			commandSelected = true
		}
	}
	if len(args) < 1 {
		log(1, "Indiepkg Version 0.1.3, run %sindiepkg help%s for usage.", textFx["BOLD"], RESETCOL)
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

const helpMsg = `Usage: indiepkg [<options>...] <command>

Commands:
  help                       Show this help message.
  install <packages...>      Installs packages.
  uninstall <packages...>    Removes packages.
  update [packages...]       Re-downloads the a package's info & install instructions. If no packages are specified, all packages are updated.
  upgrade [packages...]      Pulls git repository & recompile's a package. If no package is specified, all packages are upgraded.
  info <package>             Displays information about a specific package.
  remove-data <packages...>  Removes package data from .indiepkg. Use this only if a package installation has failed and the uninstall command won't work.
  sync                       Sync package info & package source.
  version                    Show version.

Options:
  -p, --purge                Removes a package's configuration files as well as the package itself.
  -d, --debug                Displays variable & debugging information.
  -y, --assumeyes            Assumes yes to all prompts. (Use with caution!)

Examples:
  indiepkg install my-pkg
  indiepkg uninstall other-pkg
  indiepkg upgrade third-pkg
`

const version = "0.9"

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

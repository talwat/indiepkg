package main

import (
	"os"
	"strings"
)

func parseInput() {
	args := os.Args[1:]

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") { // Flags
			switch arg {
			case "-purge":
				log(0, "purging")
				return
			}
		} else { // Commands
			switch arg {
			case "install":

			}
		}
	}
	log(0, "Indiepkg Version 0.1.3")
}

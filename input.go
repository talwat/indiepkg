package main

import (
	"os"
	"strings"
)

func parseInput() {
	args := os.Args[1:]

	for _, arg := range args {
		// Flags
		if strings.HasPrefix(arg, "-") {
			switch arg {
			case "-purge":
				log(0, "purging")
				return
			}
		}
	}
	log(0, "sus")
}

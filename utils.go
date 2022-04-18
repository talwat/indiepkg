package main

import (
	"os"
	"strings"
)

var home string = os.Getenv("HOME") + "/"

func errIs404(err error) bool {
	log(1, "Checking for a 404 error...")
	return err != nil && strings.Contains(err.Error(), "404")
}

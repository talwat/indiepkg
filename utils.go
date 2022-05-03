package main

import (
	"os"
	"os/user"
	"strings"
)

var home string = os.Getenv("HOME") + "/"

func errIs404(err error) bool {
	log(1, "Checking for a 404 error...")

	return err != nil && strings.Contains(err.Error(), "404")
}

func isRoot() bool {
	currentUser, err := user.Current()

	errorLog(err, "An error occurred while checking if user is root")

	return currentUser.Username == "root"
}

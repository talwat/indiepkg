package main

import (
	"net/url"
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

func isURL(urlInput string) bool {
	log(1, "Checking if %s is a valid URL...", bolden(urlInput))
	_, err := url.ParseRequestURI(urlInput)

	return err == nil
}

func checkFor404(statusCode int, pkgName string) bool {
	log(1, "Checking status code for 404 error...")

	switch {
	case statusCode >= 200 && statusCode <= 299 && statusCode != 204:
		return false
	case statusCode == 404 || statusCode == 204:
		return true
	default:
		errorLogRaw("An HTTP error occurred while getting package information for %s. StatusCode: %s", bolden(pkgName), bolden(statusCode))

		return false
	}
}

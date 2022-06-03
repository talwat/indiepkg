package main

import (
	"net/url"
	"os"
	"os/user"
)

var home string = os.Getenv("HOME") + "/"

// Check if user is root.
func isRoot() bool {
	currentUser, err := user.Current()

	errorLog(err, "An error occurred while checking if user is root")

	return currentUser.Username == "root"
}

// Check is URL is valid.
func isValidURL(urlInput string) bool {
	log(1, "Checking if %s is a valid URL...", bolden(urlInput))
	_, err := url.ParseRequestURI(urlInput)

	return err == nil
}

// Check for status code 404, 204, and trigger an error if not in the 200 range.
func checkFor404(statusCode int, pkgName string) bool {
	log(1, "Checking status code for 404 error...")

	switch {
	case statusCode >= 200 && statusCode <= 299 && statusCode != 204: // Check for HTTP status code to be in valid range
		return false
	case statusCode == 404 || statusCode == 204:
		return true
	default:
		errorLogRaw("An HTTP error occurred while getting package information for %s. StatusCode: %s", bolden(pkgName), bolden(statusCode))

		return false
	}
}

package main

import (
	"os"
	"os/exec"
	"strings"
)

var home string = os.Getenv("HOME") + "/"

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	errorLog(err, 4, "An error occurred while clearing screen.")
}

func errIs404(err error) bool {
	return err != nil && strings.Contains(err.Error(), "404")
}

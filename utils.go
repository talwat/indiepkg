package main

import (
	"os"
	"os/exec"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	errorLog(err, 4, "An error occurred while clearing screen.")
}

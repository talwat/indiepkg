package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func runCommand(workDir string, command string, args ...string) (string, int) {
	var cmd *exec.Cmd

	cmd = exec.Command(command, args...)
	cmd.Dir = workDir
	data, err := cmd.CombinedOutput()
	errCode := 0
	if err != nil {
		errCode, _ = strconv.Atoi(err.Error()[12:])
		log(4, "An error occurred while running command: %s %s\n    Output: %s", command, strings.Join(args, " "), string(data))
		os.Exit(1)
	}
	return strings.TrimSuffix(string(data), "\n"), errCode
}

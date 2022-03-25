package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func runCommand(workDir string, command string, args ...string) (string, error, int) {
	var cmd *exec.Cmd

	cmd = exec.Command(command, args...)
	cmd.Dir = workDir
	data, err := cmd.CombinedOutput()
	errCode := 0
	if err != nil {
		errCode, _ = strconv.Atoi(err.Error()[12:])
	}
	return strings.TrimSuffix(string(data), "\n"), err, errCode
}

package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func runCommand(workDir string, cmd string, args ...string) (string, int) {
	var cmdObj *exec.Cmd

	cmdObj = exec.Command(cmd, args...)
	cmdObj.Dir = workDir
	data, err := cmdObj.CombinedOutput()
	errCode := 0
	if err != nil {
		errCode, _ = strconv.Atoi(err.Error()[12:])
		log(4, "An error occurred while running command: %s %s\n    Output: %s", cmd, strings.Join(args, " "), string(data))
		os.Exit(1)
	}
	return strings.TrimSuffix(string(data), "\n"), errCode
}

func checkIfCommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

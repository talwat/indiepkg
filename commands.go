package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func runCommand(command string) (string, error, int) {
	splitCommand := strings.Split(command, " ")
	args := strings.Join(splitCommand[1:], " ")
	cmd := exec.Command(splitCommand[0], args)
	data, err := cmd.CombinedOutput()
	errCode := 0
	if err != nil {
		errCode, _ = strconv.Atoi(err.Error()[12:])
	}
	return string(data), err, errCode
}

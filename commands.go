package main

import (
	"bufio"
	"fmt"
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
		log(4, "An error occurred while running command: %s %s\n    Output: %s\n    Working Directory: %s\n    Error: %s", cmd, strings.Join(args, " "), string(data), workDir, err.Error())
		os.Exit(1)
	}
	return strings.TrimSuffix(string(data), "\n"), errCode
}

func runCommandRealTime(workDir string, cmd string, args ...string) {
	cmdObj := exec.Command(cmd, args...)
	cmdObj.Dir = workDir

	stdout, err := cmdObj.StdoutPipe()
	errorLogNewlineBefore(err, 4, "An error occurred while creating stdout pipe")

	err = cmdObj.Start()
	errorLogNewlineBefore(err, 4, "An error occurred while starting command")

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Printf(textCol["VIOLET"] + "." + RESETCOL)
	}

	cmdObj.Wait()
	fmt.Printf("\n")
}

func checkIfCommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

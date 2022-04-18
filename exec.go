package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func runCommandRealTime(workDir string, cmd string, args ...string) {
	cmdObj := exec.Command(cmd, args...)
	cmdObj.Dir = workDir

	cmdReader, err := cmdObj.StdoutPipe()
	cmdObj.Stderr = cmdObj.Stdout

	errorLogNewlineBefore(err, 4, "An error occurred while creating stdout pipe")

	if debug {
		fmt.Print("\n")
	}

	err = cmdObj.Start()
	errorLogNewlineBefore(err, 4, "An error occurred while starting command")

	scanner := bufio.NewScanner(cmdReader)
	for scanner.Scan() {
		if debug {
			fmt.Printf(logType[5]+(" %s\n"), scanner.Text())
		} else {
			fmt.Printf(textCol["VIOLET"] + "." + RESETCOL)
		}
	}

	err = cmdObj.Wait()
	errorLogNewlineBefore(err, 4, "An error occurred while running command to finish")
	fmt.Printf("\n")
}

func checkIfCommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

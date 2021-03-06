package main

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/talwat/spingo"
)

// Runs a command and prints the output to the terminal in real time.
func runCommandRealTime(workDir string, cmd string, args ...string) {
	cmdObj := exec.Command(cmd, args...)
	cmdObj.Dir = workDir

	// Get output pipes
	cmdReader, err := cmdObj.StdoutPipe()
	cmdObj.Stderr = cmdObj.Stdout

	errorLogNewlineBefore(err, "An error occurred while creating stdout pipe")

	if debug {
		rawLog("\n")
	}

	err = cmdObj.Start()
	errorLogNewlineBefore(err, "An error occurred while starting command")

	for {
		tmp := make([]byte, 1024)
		_, err := cmdReader.Read(tmp)
		rawLog(string(tmp))

		if err != nil {
			break
		}
	}

	err = cmdObj.Wait()
	errorLogNewlineBefore(err, "An error occurred while running command")
}

// Runs a command and displays dots to indicate progress.
func runCommandDot(commandName string, workDir string, forceCmd bool, cmd string, args ...string) {
	spinner := spingo.Spinner{ // nolint:exhaustruct
		Prefix: logType[1] + " Running command " + bolden(commandName) + " ",
		SpinnerArt: []string{
			textCol.Violet + "/" + RESETCOL,
			textCol.Violet + "|" + RESETCOL,
			textCol.Violet + "\\" + RESETCOL,
			textCol.Violet + "-" + RESETCOL,
		},
	}

	debugLog("Progress indicator: %s", config.Progress.CompileProgressIndicator)

	if config.Progress.CompileProgressIndicator == "dots" {
		logNoNewline(1, "Running command %s", bolden(commandName))
	} else {
		spinner.DisplaySpinner()
	}

	parsedCmd := strings.TrimPrefix(cmd, "!(FORCE)! ")
	cmdObj := exec.Command(parsedCmd, args...)
	cmdObj.Dir = workDir

	cmdReader, err := cmdObj.StdoutPipe()
	cmdObj.Stderr = cmdObj.Stdout

	errorLogNewlineBefore(err, "An error occurred while creating stdout pipe")

	if debug {
		rawLog("\n")
	}

	err = cmdObj.Start()
	errorLogNewlineBefore(err, "An error occurred while starting command")

	output := ""
	scanner := bufio.NewScanner(cmdReader)

	for scanner.Scan() {
		output += scanner.Text() + "\n"

		if debug {
			rawLogf(logType[5]+(" %s\n"), scanner.Text())
		} else {
			if config.Progress.CompileProgressIndicator == "dots" {
				rawLog(textCol.Violet + "." + RESETCOL)
			} else {
				spinner.DisplaySpinner()
			}
		}
	}

	output = strings.TrimSpace(output)
	err = cmdObj.Wait()

	if forceCmd {
		log(1, "Command is marked as force, so not checking for error.")
	} else {
		errorLogNewlineBefore(err, "An error occurred while running command. Output: %s", output)
	}
}

// Runs a command.
func runCommand(workDir string, cmd string, args ...string) (string, error) {
	cmdObj := exec.Command(cmd, args...)
	cmdObj.Dir = workDir

	data, err := cmdObj.CombinedOutput()
	if err != nil {
		return string(data), err
	}

	return strings.TrimSuffix(string(data), "\n"), nil
}

// Check if a command exists.
func checkIfCommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)

	return err == nil
}

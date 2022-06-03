package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ztrue/tracerr"
)

var logType = map[int]string{
	0: RESETCOL + textCol.Green + "[^]" + RESETCOL,
	1: RESETCOL + textCol.Cyan + "[.]" + RESETCOL,
	2: RESETCOL + textCol.Blue + "[#]" + RESETCOL,
	3: RESETCOL + textCol.Yellow + "[*]" + RESETCOL,
	4: RESETCOL + textCol.Red + "[!]" + RESETCOL,
	5: RESETCOL + textCol.Violet + "[+]" + RESETCOL,
	6: RESETCOL + textCol.Cyan + "[?]" + RESETCOL,
}

// Chapter log, used to organize logs into sections.
func chapLog(prefix string, colorInput string, msg string, params ...interface{}) {
	var color string

	colors := map[int]string{
		2: textCol.Violet,
		3: textCol.Blue,
		4: textCol.Cyan,
		5: textCol.Cyan,
	}

	if colorInput != "" { // Check if color is specified
		color = colorInput
	} else { // Pick color based on prefix length
		color = colors[len(prefix)]
	}

	rawLogf("\n"+RESETCOL+color+bolden(prefix+RESETCOL+textFx.Bold+" %s\n"), fmt.Sprintf(msg, params...))
}

// Adds a tab to the beginning of each line and prints it.
func indent(text string) {
	for _, line := range strings.Split(text, "\n") {
		rawLog("      " + line + "\n")
	}
}

// Logs a message and formats it.
func rawLogf(msg string, params ...interface{}) {
	fmt.Printf(msg, params...) // nolint:forbidigo
}

// Logs a message.
func rawLog(msg string) {
	fmt.Print(msg) // nolint:forbidigo
}

// Logs a message with a specified prefix.
func log(logTypeInput int, msg string, params ...interface{}) {
	rawLogf(logType[logTypeInput]+(" %s\n"), fmt.Sprintf(msg, params...))
}

// Logs a message without a newline.
func logNoNewline(logTypeInput int, msg string, params ...interface{}) {
	rawLogf(logType[logTypeInput]+(" %s"), fmt.Sprintf(msg, params...))
}

// Logs a message with an extra newline before.
func logNewlineBefore(logTypeInput int, msg string, params ...interface{}) {
	rawLogf("\n"+logType[logTypeInput]+(" %s\n"), fmt.Sprintf(msg, params...))
}

// Checks if err is not nil and if so suspend the program.
func errorLog(err error, msg string, params ...interface{}) {
	if err != nil {
		msg := fmt.Sprintf(("%s. Error: %s"), fmt.Sprintf(msg, params...), err.Error())

		if force {
			log(4, msg)
			log(3, "Continuing despite error because force is enabled...")

			return
		}

		chapLog("=>", textCol.Red, "Error")
		log(4, msg)
		log(4, "Source error log:")

		errLog := tracerr.SprintSourceColor(tracerr.Wrap(err), 6)

		for _, line := range strings.Split(errLog, "\n\n")[2:] {
			rawLog("    " + line + "\n")
		}

		os.Exit(1)
	}
}

// Suspends the program with an error message.
func errorLogRaw(msg string, params ...interface{}) {
	errMsg := fmt.Sprintf(("%s."), fmt.Sprintf(msg, params...))

	if force {
		log(4, errMsg)
		log(3, "Continuing despite error because force is enabled...")

		return
	}

	chapLog("=>", textCol.Red, "Error")
	log(4, errMsg)
	os.Exit(1)
}

// Same as errorLog but adds an extra newline before.
func errorLogNewlineBefore(err error, msg string, params ...interface{}) {
	if err != nil {
		msg := fmt.Sprintf(("%s. Error: %s"), fmt.Sprintf(msg, params...), err.Error())
		if force {
			logNewlineBefore(4, msg)
			log(3, "Continuing despite error because force is enabled...")

			return
		}

		chapLog("\n=>", textCol.Red, "Error")
		logNewlineBefore(4, msg)
		log(4, "Source error log:")

		errLog := tracerr.SprintSourceColor(tracerr.Wrap(err), 6)

		for _, line := range strings.Split(errLog, "\n")[2:] {
			rawLog("    " + line + "\n")
		}

		os.Exit(1)
	}
}

// Gets user input.
func input(defVal string, msg string, params ...interface{}) string {
	logNoNewline(6, ("%s")+": ", fmt.Sprintf(msg, params...))

	if assumeYes {
		rawLog("\n")

		return defVal
	}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

// Confirms in a y/n prompt.
func confirm(defVal, msg string, params ...interface{}) {
	if !strings.Contains(input(defVal, msg, params...), "y") {
		os.Exit(1)
	}
}

// Logs if debug is set to true.
func debugLog(msg string, params ...interface{}) {
	if debug {
		log(5, msg, params...)
	}
}

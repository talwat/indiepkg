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

func chapLog(prefix string, colorInput string, msg string, params ...interface{}) {
	var color string

	if colorInput != "" { // Check if color is specified
		color = colorInput
	} else { // Pick color based on prefix length
		switch len(prefix) {
		case 2:
			color = textCol.Violet
		case 3:
			color = textCol.Blue
		case 4:
			color = textCol.Cyan
		default:
			color = textCol.Cyan
		}
	}

	rawLogf("\n"+RESETCOL+color+bolden(prefix+RESETCOL+textFx.Bold+" %s\n"), fmt.Sprintf(msg, params...))
}

func indent(text string) {
	for _, line := range strings.Split(text, "\n") {
		rawLog("      " + line + "\n")
	}
}

func rawLogf(msg string, params ...interface{}) {
	fmt.Printf(msg, params...) // nolint:forbidigo
}

func rawLog(msg string) {
	fmt.Print(msg) // nolint:forbidigo
}

func log(logTypeInput int, msg string, params ...interface{}) {
	rawLogf(logType[logTypeInput]+(" %s\n"), fmt.Sprintf(msg, params...))
}

func logNoNewline(logTypeInput int, msg string, params ...interface{}) {
	rawLogf(logType[logTypeInput]+(" %s"), fmt.Sprintf(msg, params...))
}

func logNewlineBefore(logTypeInput int, msg string, params ...interface{}) {
	rawLogf("\n"+logType[logTypeInput]+(" %s\n"), fmt.Sprintf(msg, params...))
}

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

func confirm(defVal, msg string, params ...interface{}) {
	if !strings.Contains(input(defVal, msg, params...), "y") {
		os.Exit(1)
	}
}

func debugLog(msg string, params ...interface{}) {
	if debug {
		log(5, msg, params...)
	}
}

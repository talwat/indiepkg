package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ztrue/tracerr"
)

var logType = map[int]string{
	0: RESETCOL + textCol["GREEN"] + "[^]" + RESETCOL,
	1: RESETCOL + textCol["CYAN"] + "[.]" + RESETCOL,
	2: RESETCOL + textCol["BLUE"] + "[#]" + RESETCOL,
	3: RESETCOL + textCol["YELLOW"] + "[*]" + RESETCOL,
	4: RESETCOL + textCol["RED"] + "[!]" + RESETCOL,
	5: RESETCOL + textCol["VIOLET"] + "[+]" + RESETCOL,
	6: RESETCOL + textCol["CYAN"] + "[?]" + RESETCOL,
}

func chapLog(prefix string, colorInput string, msg string, params ...interface{}) {
	var color string = colorInput

	if colorInput != "" {
		color = colorInput
	} else {
		switch len(prefix) {
		case 2:
			color = "VIOLET"
		case 3:
			color = "BLUE"
		case 4:
			color = "CYAN"
		}
	}

	fmt.Printf("\n"+RESETCOL+textCol[color]+textFx["BOLD"]+prefix+RESETCOL+textFx["BOLD"]+(" %s\n")+RESETCOL, fmt.Sprintf(msg, params...))
}

func log(logTypeInput int, msg string, params ...interface{}) {
	fmt.Printf(logType[logTypeInput]+(" %s\n"), fmt.Sprintf(msg, params...))
}

func logNoNewline(logTypeInput int, msg string, params ...interface{}) {
	fmt.Printf(logType[logTypeInput]+(" %s"), fmt.Sprintf(msg, params...))
}

func errorLog(err error, logTypeInput int, msg string, params ...interface{}) {
	if err != nil {
		msg := fmt.Sprintf(("%s. Error: %s"), fmt.Sprintf(msg, params...), err.Error())
		if logTypeInput == 4 {
			if force {
				log(4, msg)
				log(3, "Continuing despite error because force is enabled...")
				return
			} else {
				chapLog("=>", "RED", "Error")
				log(4, msg)
				log(4, "Source error log:")
				errLog := tracerr.SprintSourceColor(tracerr.Wrap(err), 9)
				for _, line := range strings.Split(errLog, "\n")[2:] {
					fmt.Println("    " + line)
				}
			}
			os.Exit(1)
		}
	}
}

func errorLogNewlineBefore(err error, logTypeInput int, msg string, params ...interface{}) {
	if err != nil {
		fmt.Printf("\n"+logType[logTypeInput]+(" %s. Error: %s\n"), fmt.Sprintf(msg, params...), err.Error())
		if logTypeInput == 4 {
			if force {
				log(3, "Continuing despite error because force is enabled...")
				return
			}
			os.Exit(1)
		}
	}
}

func input(defVal string, msg string, params ...interface{}) string {
	if assumeYes {
		return defVal
	} else {
		reader := bufio.NewReader(os.Stdin)
		logNoNewline(6, ("%s")+": ", fmt.Sprintf(msg, params...))
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}
}

func confirm(defVal, msg string) {
	if !strings.Contains(input(defVal, msg), "y") {
		os.Exit(1)
	}
}

func debugLog(msg string, params ...interface{}) {
	if debug {
		fmt.Printf(logType[5]+(" %s\n"), fmt.Sprintf(msg, params...))
	}
}

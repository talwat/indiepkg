package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var logType = map[int]string{
	0: textCol["GREEN"] + "[!]" + RESETCOL,
	1: textCol["CYAN"] + "[!]" + RESETCOL,
	2: textCol["BLUE"] + "[!]" + RESETCOL,
	3: textCol["YELLOW"] + "[!]" + RESETCOL,
	4: textCol["RED"] + "[!]" + RESETCOL,
	5: textCol["VIOLET"] + "[+]" + RESETCOL,
}

func log(logTypeInput int, message string, params ...interface{}) {
	fmt.Printf(logType[logTypeInput]+(" %s\n"), fmt.Sprintf(message, params...))
}

func logNoNewline(logTypeInput int, message string, params ...interface{}) {
	fmt.Printf(logType[logTypeInput]+(" %s"), fmt.Sprintf(message, params...))
}

func errorLog(err error, logTypeInput int, message string, params ...interface{}) {
	if err != nil {
		fmt.Printf(logType[logTypeInput]+(" %s. Error: %s\n"), fmt.Sprintf(message, params...), err.Error())
		if logTypeInput == 4 {
			if force {
				log(3, "Continuing despite error because force is enabled...")
				return
			}
			os.Exit(1)
		}
	}
}

func errorLogNewlineBefore(err error, logTypeInput int, message string, params ...interface{}) {
	if err != nil {
		fmt.Printf("\n"+logType[logTypeInput]+(" %s. Error: %s\n"), fmt.Sprintf(message, params...), err.Error())
		if logTypeInput == 4 {
			if force {
				log(3, "Continuing despite error because force is enabled...")
				return
			}
			os.Exit(1)
		}
	}
}

func input(defVal string, message string, params ...interface{}) string {
	if assumeYes {
		return defVal
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf(textCol["CYAN"]+"[?]"+RESETCOL+(" %s")+": ", fmt.Sprintf(message, params...))
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}
}

func confirm(defVal, message string) {
	if !strings.Contains(input(defVal, message), "y") {
		os.Exit(1)
	}
}

func debugLog(message string, params ...interface{}) {
	if debug {
		fmt.Printf(logType[5]+(" %s\n"), fmt.Sprintf(message, params...))
	}
}

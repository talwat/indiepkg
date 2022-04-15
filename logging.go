package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var logType = map[int]string{
	0: RESETCOL + textCol["GREEN"] + "[^]" + RESETCOL,
	1: RESETCOL + textCol["CYAN"] + "[!]" + RESETCOL,
	2: RESETCOL + textCol["BLUE"] + "[!]" + RESETCOL,
	3: RESETCOL + textCol["YELLOW"] + "[!]" + RESETCOL,
	4: RESETCOL + textCol["RED"] + "[!]" + RESETCOL,
	5: RESETCOL + textCol["VIOLET"] + "[+]" + RESETCOL,
	6: RESETCOL + textCol["CYAN"] + "[?]" + RESETCOL,
}

func chapLog(prefix string, color string, message string, params ...interface{}) {
	fmt.Printf("\n"+RESETCOL+textCol[color]+textFx["BOLD"]+prefix+textCol["WHITE"]+(" %s\n")+RESETCOL, fmt.Sprintf(message, params...))
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
		fmt.Printf(logType[6]+(" %s")+": ", fmt.Sprintf(message, params...))
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

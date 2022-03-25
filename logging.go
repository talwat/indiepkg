package main

import (
	"bufio"
	"fmt"
	"os"
)

var logType = map[int]string{
	0: textCol["GREEN"] + "[!]" + RESETCOL,
	1: textCol["CYAN"] + "[!]" + RESETCOL,
	2: textCol["BLUE"] + "[!]" + RESETCOL,
	3: textCol["YELLOW"] + "[!]" + RESETCOL,
	4: textCol["RED"] + "[!]" + RESETCOL,
}

func log(logTypeInput int, message string, params ...interface{}) {
	fmt.Printf(logType[logTypeInput]+(" %s\n"), fmt.Sprintf(message, params...))
}

func errorLog(err error, logTypeInput int, message string, params ...interface{}) {
	if err != nil {
		fmt.Printf(logType[logTypeInput]+(" %s Error: %s\n"), fmt.Sprintf(message, params...), err.Error())
		if logTypeInput == 4 {
			os.Exit(1)
		}
	}
}

func input(message string, params ...interface{}) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(textCol["CYAN"]+"[!]"+RESETCOL+(" %s")+": ", fmt.Sprintf(message, params...))
	input, _ := reader.ReadString('\n')
	return input
}

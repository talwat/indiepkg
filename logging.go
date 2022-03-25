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

func input(message string, params ...interface{}) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(textCol["CYAN"]+"[!]"+RESETCOL+(" %s")+": ", fmt.Sprintf(message, params...))
	input, _ := reader.ReadString('\n')
	return input
}

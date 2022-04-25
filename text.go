package main

var RESETCOL string = "\x1B[0m"

var textFx = map[string]string{
	"END":           "\x1B[0m",
	"BOLD":          "\x1B[1m",
	"DIM":           "\x1B[m",
	"ITALIC":        "\x1B[3m",
	"URL":           "\x1B[4m",
	"BLINK":         "\x1B[5m",
	"BLINK2":        "\x1B[6m",
	"SELECTED":      "\x1B[7m",
	"STRIKETHROUGH": "\x1B[9m",
}

var textCol = map[string]string{
	"BLACK":  "\x1B[30m",
	"RED":    "\x1B[31m",
	"GREEN":  "\x1B[32m",
	"YELLOW": "\x1B[33m",
	"BLUE":   "\x1B[34m",
	"VIOLET": "\x1B[35m",
	"CYAN":   "\x1B[36m",
	"WHITE":  "\x1B[37m",
}

var textBgCol = map[string]string{
	"BLACK":  "\x1B[40m",
	"RED":    "\x1B[41m",
	"GREEN":  "\x1B[4m",
	"YELLOW": "\x1B[43m",
	"BLUE":   "\x1B[44m",
	"VIOLET": "\x1B[45m",
	"BEIGE":  "\x1B[46m",
	"WHITE":  "\x1B[47m",
}

var lightTextCol = map[string]string{
	"BLACK":  "\x1B[90m",
	"RED":    "\x1B[91m",
	"GREEN":  "\x1B[9m",
	"YELLOW": "\x1B[93m",
	"BLUE":   "\x1B[94m",
	"VIOLET": "\x1B[95m",
	"BEIGE":  "\x1B[96m",
	"WHITE":  "\x1B[97m",
}

var lightTextBgCol = map[string]string{
	"GREY":   "\x1B[100m",
	"RED":    "\x1B[101m",
	"GREEN":  "\x1B[10m",
	"YELLOW": "\x1B[103m",
	"BLUE":   "\x1B[104m",
	"VIOLET": "\x1B[105m",
	"BEIGE":  "\x1B[106m",
	"WHITE":  "\x1B[107m",
}

func bolden(text string) string {
	return textFx["BOLD"] + text + RESETCOL
}

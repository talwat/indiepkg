package main

import "fmt"

type Color struct {
	Black  string
	Red    string
	Green  string
	Yellow string
	Blue   string
	Violet string
	Cyan   string
	Beige  string
	White  string
}

type Effects struct {
	End           string
	Bold          string
	Dim           string
	Italic        string
	URL           string
	Blink         string
	Blink2        string
	Selected      string
	Strikethrough string
}

var RESETCOL = "\x1B[0m"

var textFx = Effects{
	End:           "\x1B[0m",
	Bold:          "\x1B[1m",
	Dim:           "\x1B[m",
	Italic:        "\x1B[3m",
	URL:           "\x1B[4m",
	Blink:         "\x1B[5m",
	Blink2:        "\x1B[6m",
	Selected:      "\x1B[7m",
	Strikethrough: "\x1B[9m",
}

var textCol = Color{
	Black:  "\x1B[30m",
	Red:    "\x1B[31m",
	Green:  "\x1B[32m",
	Yellow: "\x1B[33m",
	Blue:   "\x1B[34m",
	Violet: "\x1B[35m",
	Cyan:   "\x1B[36m",
	White:  "\x1B[37m",
	Beige:  "\x1B[37m",
}

var textBgCol = Color{
	Black:  "\x1B[40m",
	Red:    "\x1B[41m",
	Green:  "\x1B[4m",
	Yellow: "\x1B[43m",
	Blue:   "\x1B[44m",
	Cyan:   "\x1B[44m",
	Violet: "\x1B[45m",
	Beige:  "\x1B[46m",
	White:  "\x1B[47m",
}

var lightTextCol = Color{
	Black:  "\x1B[90m",
	Red:    "\x1B[91m",
	Green:  "\x1B[9m",
	Yellow: "\x1B[93m",
	Blue:   "\x1B[94m",
	Cyan:   "\x1B[94m",
	Violet: "\x1B[95m",
	Beige:  "\x1B[96m",
	White:  "\x1B[97m",
}

var lightTextBgCol = Color{
	Black:  "\x1B[100m",
	Red:    "\x1B[101m",
	Green:  "\x1B[10m",
	Yellow: "\x1B[103m",
	Blue:   "\x1B[104m",
	Cyan:   "\x1B[104m",
	Violet: "\x1B[105m",
	Beige:  "\x1B[106m",
	White:  "\x1B[107m",
}

func bolden(text interface{}) string {
	return textFx.Bold + fmt.Sprint(text) + RESETCOL
}

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

var RESETCOL = "\x1B[0"

var textFx = Effects{
	End:           "\x1B[0",
	Bold:          "\x1B[1",
	Dim:           "\x1B[",
	Italic:        "\x1B[3",
	URL:           "\x1B[4",
	Blink:         "\x1B[5",
	Blink2:        "\x1B[6",
	Selected:      "\x1B[7",
	Strikethrough: "\x1B[9",
}

var textCol = Color{
	Black:  "\x1B[30",
	Red:    "\x1B[31",
	Green:  "\x1B[32",
	Yellow: "\x1B[33",
	Blue:   "\x1B[34",
	Violet: "\x1B[35",
	Cyan:   "\x1B[36",
	White:  "\x1B[37",
}

var textBgCol = Color{
	Black:  "\x1B[40",
	Red:    "\x1B[41",
	Green:  "\x1B[4",
	Yellow: "\x1B[43",
	Blue:   "\x1B[44",
	Violet: "\x1B[45",
	Beige:  "\x1B[46",
	White:  "\x1B[47",
}

var lightTextCol = Color{
	Black:  "\x1B[90",
	Red:    "\x1B[91",
	Green:  "\x1B[9",
	Yellow: "\x1B[93",
	Blue:   "\x1B[94",
	Violet: "\x1B[95",
	Beige:  "\x1B[96",
	White:  "\x1B[97",
}

var lightTextBgCol = Color{
	Black:  "\x1B[100",
	Red:    "\x1B[101",
	Green:  "\x1B[10",
	Yellow: "\x1B[103",
	Blue:   "\x1B[104",
	Violet: "\x1B[105",
	Beige:  "\x1B[106",
	White:  "\x1B[107",
}

func bolden(text interface{}) string {
	return textFx.Bold + fmt.Sprint(text) + RESETCOL
}

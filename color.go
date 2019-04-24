package log

import (
	"strconv"
)

const (
	// common
	Reset  = "\033[0m" // auto reset the rest of text to default color
	Normal = 0
	Bold   = 1 // increase this value if you want bolder text
	// special
	Dim       = 2
	Underline = 4
	Blink     = 5
	Reverse   = 7
	Hidden    = 8
	// color
	Black       = 30 // default = 39
	Red         = 31
	Green       = 32
	Yellow      = 33
	Blue        = 34
	Purple      = 35 // purple = magenta
	Cyan        = 36
	LightGray   = 37
	DarkGray    = 90
	LightRed    = 91
	LightGreen  = 92
	LightYellow = 93
	LightBlue   = 94
	LightPurple = 95
	LightCyan   = 96
	White       = 97
)

// you can use custom color code and font size by calling this function
func Render(colorCode int, fontSize int, content string) string {
	return "\033[" + strconv.Itoa(fontSize) + ";" + strconv.Itoa(colorCode) + "m" + content + Reset
}

// black text (use this with caution since most geeks use dark console)
func Test(txt string) string {
	return Render(Black, Normal, txt)
}

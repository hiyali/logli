package logli

import (
	"strconv"
)

const (
	reset = "\033[0m" // auto reset the rest of text to default color
	// type face
	normal = 0
	bold   = 1 // increase this value if you want bolder text
	// cursor
	dim       = 2
	underline = 4
	blink     = 5
	reverse   = 7
	hidden    = 8
	// color
	black       = 30 // default = 39
	red         = 31
	green       = 32
	yellow      = 33
	blue        = 34
	purple      = 35 // purple = magenta
	cyan        = 36
	lightGray   = 37
	darkGray    = 90
	lightRed    = 91
	lightGreen  = 92
	lightYellow = 93
	lightBlue   = 94
	lightPurple = 95
	lightCyan   = 96
	white       = 97
	_default    = -1
)

type (
	typeFaceType struct {
		Normal int
		Bold   int
	}

	cursorType struct {
		Dim       int
		Underline int
		Blink     int
		Reverse   int
		Hidden    int
	}

	color     int
	colorType struct {
		Black       color
		Red         color
		Green       color
		Yellow      color
		Blue        color
		Purple      color
		Cyan        color
		LightGray   color
		DarkGray    color
		LightRed    color
		LightGreen  color
		LightYellow color
		LightBlue   color
		LightPurple color
		LightCyan   color
		White       color
		Default     color
	}
)

var (
	Color    colorType
	Cursor   cursorType
	TypeFace typeFaceType
)

func (*colorType) Print() {
	Debug("Color.")
	PrintF("\tBlack - %v", Colorize("Black", Color.Black))
	PrintF("\tRed - %v", Colorize("Red", Color.Red))
	PrintF("\tGreen - %v", Colorize("Green", Color.Green))
	PrintF("\tYellow - %v", Colorize("Yellow", Color.Yellow))
	PrintF("\tBlue - %v", Colorize("Blue", Color.Blue))
	PrintF("\tPurple - %v", Colorize("Purple", Color.Purple))
	PrintF("\tCyan - %v", Colorize("Cyan", Color.Cyan))
	PrintF("\tLightGray - %v", Colorize("LightGray", Color.LightGray))
	PrintF("\tDarkGray - %v", Colorize("DarkGray", Color.DarkGray))
	PrintF("\tLightRed - %v", Colorize("LightRed", Color.LightRed))
	PrintF("\tLightGreen - %v", Colorize("LightGreen", Color.LightGreen))
	PrintF("\tLightYellow - %v", Colorize("LightYellow", Color.LightYellow))
	PrintF("\tLightBlue - %v", Colorize("LightBlue", Color.LightBlue))
	PrintF("\tLightPurple - %v", Colorize("LightPurple", Color.LightPurple))
	PrintF("\tLightCyan - %v", Colorize("LightCyan", Color.LightCyan))
	PrintF("\tWhite - %v", Colorize("White", Color.White))
	PrintF("\tDefault - %v", Colorize("Default", Color.Default))
}

func (*cursorType) Print() {
	Debug("Cursor.")
	CustomPrint("\tDim", Color.Default, Cursor.Dim)
	CustomPrint("\tUnderline", Color.Default, Cursor.Underline)
	CustomPrint("\tBlink", Color.Default, Cursor.Blink)
	CustomPrint("\tReverse", Color.Default, Cursor.Reverse)
	CustomPrint("\tHidden", Color.Default, Cursor.Hidden)
}

func (*typeFaceType) Print() {
	Debug("TypeFace.")
	CustomPrint("\tNormal", Color.Default, TypeFace.Normal)
	CustomPrint("\tBold", Color.Default, TypeFace.Bold)
}

func init() {
	Color = colorType{
		black,
		red,
		green,
		yellow,
		blue,
		purple,
		cyan,
		lightGray,
		darkGray,
		lightRed,
		lightGreen,
		lightYellow,
		lightBlue,
		lightPurple,
		lightCyan,
		white,
		_default,
	}
	TypeFace = typeFaceType{
		normal,
		bold,
	}
	Cursor = cursorType{
		dim,
		underline,
		blink,
		reverse,
		hidden,
	}
}

func Colorize(msg string, color color) (result string) {
	result = vary(color, TypeFace.Normal, msg)
	return
}

// You can use custom color code and font size by calling this function
func vary(color color, typeFace int, content string) string {
	return "\033[" + strconv.Itoa(typeFace) + ";" + strconv.Itoa(int(color)) + "m" + content + reset
}

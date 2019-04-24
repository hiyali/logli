package log

import (
	"fmt"
	"os"
)

func Colorful(msg string, color int) string {
	return Render(color, Normal, msg)
}

func Special(msg string, color int, typ int) (result string) {
	result = Render(color, typ, msg)
	return
}

func Custom(msg string, color int, typ int) {
	msg = Special(msg, color, typ)
	Print(msg)
}

func Print(msg string) {
	if timeDisabled == false {
		msg = fmt.Sprintf("%v\t%v", Now(), msg)
	}
	fmt.Printf("%v\n", msg)
}

func PrintF(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if timeDisabled == false {
		msg = fmt.Sprintf("%v\t%v", Now(), msg)
	}
	fmt.Printf("%v\n", msg)
}

// Debug same to Print with color Cyan
func Debug(msg string) {
	if colorDisabled == false {
		msg = Special(msg, Cyan, Normal)
	}
	Print(msg)
}

// DebugF same to PrintF with color Cyan
func DebugF(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if colorDisabled == false {
		msg = Special(msg, Cyan, Normal)
	}
	Print(msg)
}

// Info same to Print with color Green
func Info(msg string) {
	if colorDisabled == false {
		msg = Special(msg, Green, Normal)
	}
	Print(msg)
}

// InfoF same to PrintF with color Green
func InfoF(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if colorDisabled == false {
		msg = Special(msg, Green, Normal)
	}
	Print(msg)
}

// Warn same to Print with color Yellow
func Warn(msg string) {
	if colorDisabled == false {
		msg = Special(msg, Yellow, Normal)
	}
	Print(msg)
}

// WarnF same to PrintF with color Yellow
func WarnF(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if colorDisabled == false {
		msg = Special(msg, Yellow, Normal)
	}
	Print(msg)
}

// Error same to Print with color Red
func Error(msg string) {
	if colorDisabled == false {
		msg = Special(msg, Red, Normal)
	}
	Print(msg)
}

// ErrorF same to Printf with color Red
func ErrorF(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if colorDisabled == false {
		msg = Special(msg, Red, Normal)
	}
	Print(msg)
}

// Fatal same to Error and following os.Exit(1)
func Fatal(msg string) {
	Error(msg)
	os.Exit(1)
}

// FatalF same to ErrorF and following os.Exit(1)
func FatalF(format string, a ...interface{}) {
	ErrorF(format, a...)
	os.Exit(1)
}

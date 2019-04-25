package logli

import (
	"fmt"
	"time"
)

// https://golang.org/pkg/time/
const (
	_Default = "2006-01-02 15:04:05"
	// Standard
	_ANSIC       = "Mon Jan _2 15:04:05 2006"
	_UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	_RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	_RFC822      = "02 Jan 06 15:04 MST"
	_RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	_RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	_RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	_RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	_RFC3339     = "2006-01-02T15:04:05Z07:00"
	_RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	_Kitchen     = "3:04PM"
	// Handy time stamps.
	_Stamp      = "Jan _2 15:04:05"
	_StampMilli = "Jan _2 15:04:05.000"
	_StampMicro = "Jan _2 15:04:05.000000"
	_StampNano  = "Jan _2 15:04:05.000000000"
)

type (
	time_format    string
	timeFormatType struct {
		Default time_format
		// Standard
		ANSIC       time_format
		UnixDate    time_format
		RubyDate    time_format
		RFC822      time_format
		RFC822Z     time_format
		RFC850      time_format
		RFC1123     time_format
		RFC1123Z    time_format
		RFC3339     time_format
		RFC3339Nano time_format
		Kitchen     time_format
		// Handy time stamps.
		Stamp      time_format
		StampMilli time_format
		StampMicro time_format
		StampNano  time_format
	}
)

var TimeFormat timeFormatType

func timePrint(msg string) {
	tabs := "\t"
	if len(msg) < 8 {
		tabs = "\t\t"
	}
	if timeDisabled == false {
		msg = fmt.Sprintf("%s%s%v", msg, tabs, Now())
	}
	fmt.Printf("%v\n", msg)
}

func (*timeFormatType) Print() {
	var timeFormatBefore = timeFormat

	Debug("TimeFormat.")
	timeFormat = TimeFormat.Default
	timePrint("Default")
	// -- Standard --
	timeFormat = TimeFormat.ANSIC
	timePrint("ANSIC")
	timeFormat = TimeFormat.UnixDate
	timePrint("UnixDate")
	timeFormat = TimeFormat.RubyDate
	timePrint("RubyDate")
	timeFormat = TimeFormat.RFC822
	timePrint("RFC822")
	timeFormat = TimeFormat.RFC822Z
	timePrint("RFC822Z")
	timeFormat = TimeFormat.RFC850
	timePrint("RFC850")
	timeFormat = TimeFormat.RFC1123
	timePrint("RFC1123")
	timeFormat = TimeFormat.RFC1123Z
	timePrint("RFC1123Z")
	timeFormat = TimeFormat.RFC3339
	timePrint("RFC3339")
	timeFormat = TimeFormat.RFC3339Nano
	timePrint("RFC3339Nano")
	timeFormat = TimeFormat.Kitchen
	timePrint("Kitchen")
	// -- Handy time stamps --
	timeFormat = TimeFormat.Stamp
	timePrint("Stamp")
	timeFormat = TimeFormat.StampMilli
	timePrint("StampMilli")
	timeFormat = TimeFormat.StampMicro
	timePrint("StampMicro")
	timeFormat = TimeFormat.StampNano
	timePrint("StampNano")

	timeFormat = timeFormatBefore
}

func init() {
	TimeFormat = timeFormatType{
		_Default,
		// Standard
		_ANSIC,
		_UnixDate,
		_RubyDate,
		_RFC822,
		_RFC822Z,
		_RFC850,
		_RFC1123,
		_RFC1123Z,
		_RFC3339,
		_RFC3339Nano,
		_Kitchen,
		// Handy time stamps.
		_Stamp,
		_StampMilli,
		_StampMicro,
		_StampNano,
	}
}

func SetCustomTimeFormat(format string) {
	timeFormat = time_format(format)
}

func Now() (now string) {
	now = time.Now().Format(string(timeFormat))
	if timeColor != Color.Default {
		now = Colorize(now, timeColor)
	}
	return
}

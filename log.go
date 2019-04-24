package log

import ()

var (
	timeDisabled  = false
	colorDisabled = false
	// coloredTime   = false
)

type Config struct {
	TimeDisabled  bool
	ColorDisabled bool
	// ColoredTime   bool
}

func SetConfig(conf *Config) {
	timeDisabled = conf.TimeDisabled
	colorDisabled = conf.TimeDisabled
	// coloredTime = conf.ColoredTime
}

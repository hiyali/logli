package log

func CustomTest() {
	Debug("This is a debug")
	DebugF("This is a debug that %v", "Formatted")

	Info("This is a green info")
	InfoF("This is a green info that %v", "Formatted")

	Warn("This is a warn")
	WarnF("This is a warn that %v", "Formatted")

	// SetConfig(&Config{ColoredTime: true})

	Error("This is a error")
	ErrorF("This is a error that %v", "Formatted")

	Fatal("This is a last fatal test.")
}

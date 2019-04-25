package logli

var (
	timeDisabled  = false
	colorDisabled = false
	timeColor     = color(_default)
	timeFormat    = time_format(_Default)
)

type Config struct {
	TimeDisabled  bool
	ColorDisabled bool
	TimeColor     color
	TimeFormat    time_format
}

func SetConfig(conf *Config) {
	// PrintF("SetConfig - conf: %v", conf)
	timeDisabled = conf.TimeDisabled
	colorDisabled = conf.TimeDisabled
	if conf.TimeColor > 0 {
		timeColor = conf.TimeColor
	}
	if len(conf.TimeFormat) > 0 {
		timeFormat = conf.TimeFormat
	}
}

func FullLook() {
	SetConfig(&Config{TimeDisabled: true})
	Print("SetConfig(&Config{ TimeDisabled: true })")

	Print("\nColor.Print()")
	Color.Print()
	Print("\nCursor.Print()")
	Cursor.Print()
	Print("\nTypeFace.Print()")
	TypeFace.Print()

	SetConfig(&Config{TimeDisabled: false})
	Print("SetConfig(&Config{ TimeDisabled: false })")

	CustomPrintNT("\nTimeFormat.Print()", Color.Red, TypeFace.Bold)
	TimeFormat.Print()

	Print("SetConfig(&Config{ TimeDisabled: false })")
	CustomPrintNT("\nTakeLook()", Color.Green, TypeFace.Bold)
	TakeLook()
}

func TakeLook() {
	Print("Print(msg) -> This is a default Print")
	PrintF("PrintF(format, values...) -> This is a default PrintF that %v", "Formatted")

	Debug("Debug(msg) -> This is a cyan Debug")
	DebugF("DebugF(format, values...) -> This is a cyan DebugF that %v", "Formatted")

	SetConfig(&Config{TimeColor: Color.LightPurple})
	Print("SetConfig(&Config{TimeColor: Color.LightPurple})")

	Info("Info(msg) -> This is a green Info")
	InfoF("InfoF(format, values...) -> This is a green InfoF that %v", "Formatted")

	Warn("Warn(msg) -> This is a yellow Warn")
	WarnF("WarnF(format, values...) -> This is a yellow WarnF that %v", "Formatted")

	SetConfig(&Config{TimeFormat: TimeFormat.StampNano, TimeColor: Color.LightYellow})
	Print("SetConfig(&Config{TimeFormat: TimeFormat.StampNano, TimeColor: Color.LightYellow})")

	Error("Error(msg) -> This is a red Error")
	ErrorF("ErrorF(format, values...) -> This is a red ErrorF that %v", "Formatted")

	Fatal("Fatal(msg) -> This is a last Fatal")
	FatalF("FatalF(format, values...) -> This is a FatalF that %v and %s", "Formatted", "you will never seen as above log.")
}

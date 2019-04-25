# 🏳️‍🌈 logli
> A colored logger of golang

## Take a look

#### Install
```shell
go get -u github.com/hiyali/logli
```

#### Usage
```golang
import (
  log "github.com/hiyali/logli"
)

func main() {
  log.Debug("This is a debug log actually")
  log.FatalF("Damn, there is an error: %v", err) // followed an os.Exit(1)
}
```

#### Screenshots
![Methods](https://raw.githubusercontent.com/hiyali/logli/master/screenshots/methods.png "methods")

## Methods

#### General
| Name    | Color   | Usage   |
|---      |---      |---      |
| Print   | Default | `log.Print(msg string)`                             |
| PrintF  | Default | `log.Printf(format string, values ...interface{})`  |
| Debug   | Cyan    | Same as Print, like `log.Debug(msg string)`                             |
| DebugF  | Cyan    | Same as PrintF, like `log.DebugF(format string, values ...interface{})`  |
| Info    | Green   | Same to Debug  |
| InfoF   | Green   | Same to DebugF |
| Warn    | Yellow  | Same to Debug  |
| WarnF   | Yellow  | Same to DebugF |
| Error   | Red     | Same to Debug  |
| ErrorF  | Red     | Same to DebugF |
| Fatal   | Red     | Same to Error and following os.Exit(1)    |
| FatalF  | Red     | Same to ErrorF and following os.Exit(1)   |

#### Customize
| Name          | Color       | Usage   |
|---            |---          |---      |
| CustomPrint   | As you wish | `log.CustomPrint(msg string, color color, typ int)`    |
| CustomPrintNT | Like above  | `log.CustomPrintNT(msg string, color color, typ int)` and no time will automatically print     |
| Colorize      | You will choose | `log.Colorize(msg string, color color) -> (result string)` no any print here, but would return colorized result string |

#### Other
| Name          			| Usage   |
|---            			|---      |
| SetCustomTimeFormat | `log.SetCustomTimeFormat("2006-01-02 15:04:05")` 				|
| Now									| `log.Now() -> (now string)` return time string for now 	|

## Configs

#### log.Config Options
| Option            | Type        | Default |
|---                |---          |---      |
| TimeDisabled      | bool        | false   |
| ColorDisabled     | bool        | false   |
| TimeColor         | color       | Color.Default       |
| TimeFormat        | time_format | TimeFormat.Default ("2006-01-02 15:04:05") |

> How can I set config?

```golang
log.SetConfig(&log.Config{
  TimeDisabled:   true,
  ColorDisabled:  false,
  ...
})
```

![Time Disabled True](https://raw.githubusercontent.com/hiyali/logli/master/screenshots/time_disabled_true.png "time disabled true")
![Time Disabled False](https://raw.githubusercontent.com/hiyali/logli/master/screenshots/time_disabled_false.png "time disabled false")

#### log.Color color
![Colors](https://raw.githubusercontent.com/hiyali/logli/master/screenshots/colors.png "colors")

> How can I use these colors?

```golang
log.Color.Blue -> color
...
log.CustomPrint("Hey yo", log.Color.LightRed, log.TypeFace.Normal)
```

#### log.TimeFormat time_format
![TimeFormats](https://raw.githubusercontent.com/hiyali/logli/master/screenshots/time_formats.png "time formats")

> How can I get and use these time formats?

```golang
log.TimeFormat.UnixDate -> time_format
...
log.SetConfig(&log.Config{
  TimeFormat: log.TimeFormat.Stamps,
  ...
})
```

## Print

> Print all things it can do.

```golang
log.FullLook()
```
> OR some of them
```golang
log.TakeLook()
log.Color.Print()
log.TimeFormat.Print()
```

## Contribute
> Feel free

## License
MIT

---

inspired By [bclicn/color](https://github.com/bclicn/color)

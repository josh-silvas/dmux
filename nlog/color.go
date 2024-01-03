package nlog

import (
	"fmt"
	"strconv"
)

const (
	timeFormat = "[15:04:05]"

	reset = "\033[0m"

	cyan         = 36
	lightGray    = 37
	darkGray     = 90
	lightRed     = 91
	lightYellow  = 93
	lightBlue    = 94
	lightMagenta = 95
	white        = 97
	fatalColor   = 41 // Red Background
)

// color : add color to a log message
func color(code int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(code), v, reset)
}

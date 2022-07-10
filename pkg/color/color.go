package color

import (
	"github.com/mgutz/ansi"
)

var (
	magenta = ansi.ColorFunc("magenta")
	red     = ansi.ColorFunc("red")
	redBold = ansi.ColorFunc("red+b")
	gray    = ansi.ColorFunc("black+h")
	bold    = ansi.ColorFunc("default+b")
)

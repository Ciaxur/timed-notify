package color

import (
	"github.com/mgutz/ansi"
)

var (
	Bold    = ansi.ColorFunc("default+b")
	Gray    = ansi.ColorFunc("black+h")
	Green   = ansi.ColorFunc("green")
	Magenta = ansi.ColorFunc("magenta")
	Red     = ansi.ColorFunc("red")
	RedBold = ansi.ColorFunc("red+b")
)

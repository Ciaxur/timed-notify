package Config

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

// CONFIGURE GLOBAL STD OUTPUT COLORS
var (
	ErrOut     = color.New(color.FgRed).Add(color.Bold)
	InfoOut    = color.New(color.FgHiMagenta)
	StdOut     = color.New()
	BinPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	PidDir     = "/tmp/timed-notify.pids"
	ResPath    = "/usr/share/timed-notify"
	VERSION    = "1.1.0"
)

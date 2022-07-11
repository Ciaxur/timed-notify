package config

import (
	"os"
	"path/filepath"
)

var (
	BinPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	PidDir     = "/tmp/timed-notify.pids"
	ResPath    = "/usr/share/timed-notify" // Linux Default (Adjusted in Parse.go)
)

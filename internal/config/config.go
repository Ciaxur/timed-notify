package config

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	BinPath, _       = filepath.Abs(filepath.Dir(os.Args[0]))
	PidDir           = "/tmp/timed-notify.pids"
	UserResourcePath = "/usr/share/timed-notify" // Linux Default
)

func init() {
	if runtime.GOOS == "darwin" { // MacOS Default Shared Path
		UserResourcePath = "/usr/local/share/timed-notify"
	}
}

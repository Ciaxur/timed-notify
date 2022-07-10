package build

import "time"

// The following is dynamically modified from the Makefile.
var (
	stdDateFmt = "2006-01-02"
	Version    = "dev"
	BuildDate  = time.Now().Format(stdDateFmt) // YYYY-MM-DD
)

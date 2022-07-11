package main

import (
	"fmt"
	"os"

	"github.com/ciaxur/timed-notify/v1/pkg/cmd/root"
	"github.com/ciaxur/timed-notify/v1/pkg/color"
	"github.com/ciaxur/timed-notify/v1/pkg/iostreams"
)

func main() {
	ioStream := iostreams.System()
	rootCmd := root.NewRootCmd(ioStream)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(ioStream.Err, color.Red("%v\n"), err)
		os.Exit(1)
	}
	os.Exit(0)
}

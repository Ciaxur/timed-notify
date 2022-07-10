package version

import (
	"fmt"

	"github.com/ciaxur/timed-notify/v1/internal/build"
	"github.com/ciaxur/timed-notify/v1/pkg/iostreams"
	"github.com/spf13/cobra"
)

func NewVersionCmd(ioStream *iostreams.IOStreams) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "App version",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintf(ioStream.Out, "Version: %s \nBuild Date: %s\n", build.Version, build.BuildDate)
			return err
		},
	}
	return versionCmd
}

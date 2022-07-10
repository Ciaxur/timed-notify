package root

import (
	versionCmd "github.com/ciaxur/timed-notify/v1/pkg/cmd/version"
	"github.com/ciaxur/timed-notify/v1/pkg/iostreams"
	"github.com/spf13/cobra"
)

func NewRootCmd(ioStream *iostreams.IOStreams) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "timed-notify <command> [OPTIONS]",
		Short: "Timed Notify",
		Long:  "Issue and manage timed desktop notifications",
	}
	rootCmd.AddCommand(versionCmd.NewVersionCmd(ioStream))
	return rootCmd
}

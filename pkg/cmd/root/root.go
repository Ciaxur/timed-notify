package root

import (
	"fmt"

	"github.com/ciaxur/timed-notify/v1/internal/config"
	createCmd "github.com/ciaxur/timed-notify/v1/pkg/cmd/create"
	killCmd "github.com/ciaxur/timed-notify/v1/pkg/cmd/kill"
	listCmd "github.com/ciaxur/timed-notify/v1/pkg/cmd/list"
	versionCmd "github.com/ciaxur/timed-notify/v1/pkg/cmd/version"
	"github.com/ciaxur/timed-notify/v1/pkg/iostreams"
	"github.com/spf13/cobra"
)

func NewRootCmd(ioStream *iostreams.IOStreams) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "timed-notify",
		Short: "Timed Notify",
		Long:  "Issue and manage timed desktop notifications",
	}

	// Root sub-commands.
	rootCmd.AddCommand(versionCmd.NewVersionCmd(ioStream))
	rootCmd.AddCommand(listCmd.NewListCommand(ioStream))
	rootCmd.AddCommand(createCmd.NewCreateCommand(ioStream))
	rootCmd.AddCommand(killCmd.NewKillCommand(ioStream))

	// Root flags.
	rootCmd.PersistentFlags().Bool("debug", false, fmt.Sprintf("Enables daemonized log output to %s.", config.BinPath))

	return rootCmd
}

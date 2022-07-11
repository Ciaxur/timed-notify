package kill

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ciaxur/timed-notify/v1/pkg/iostreams"
	"github.com/ciaxur/timed-notify/v1/pkg/notification"
	"github.com/spf13/cobra"
)

var killPid int

func handleKillCommand(ioStream *iostreams.IOStreams, cmd *cobra.Command, args []string) error {
	processManager, err := notification.NewNotificationManager()
	if err != nil {
		fmt.Errorf("failed to instantiate notification manager: %v", err)
	}

	// Get the process matching the pid.
	process := processManager.GetPid(killPid)
	if process != nil {
		sysProc, err := os.FindProcess(process.Pid)
		if err != nil {
			return fmt.Errorf("failed to find pid '%d' on host machine: %v", killPid, err)
		}

		if err := sysProc.Signal(os.Interrupt); err != nil {
			return fmt.Errorf("failed to interrupt process: %v", err)
		}

		fmt.Fprintf(ioStream.Out, "Pid %d interrupted.\n", process.Pid)
		return nil
	}
	return fmt.Errorf("pid %d could not be found", killPid)
}

func NewKillCommand(ioStream *iostreams.IOStreams) *cobra.Command {
	killCmd := &cobra.Command{
		Use:   "kill <PID>",
		Short: "Terminates the given Pid",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("pid argument required")
			}

			// Attempt to parse passed in pid.
			pidInt, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("failed to parse pid: %v", err)
			}

			// Validate the pid value.
			if pidInt < 1 {
				return fmt.Errorf("pid cannot be less than 0")
			}

			killPid = pidInt
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleKillCommand(ioStream, cmd, args)
		},
	}

	return killCmd
}

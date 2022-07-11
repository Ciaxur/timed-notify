package list

import (
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/ciaxur/timed-notify/v1/pkg/color"
	"github.com/ciaxur/timed-notify/v1/pkg/iostreams"
	"github.com/ciaxur/timed-notify/v1/pkg/notification"
	"github.com/spf13/cobra"
)

func handleListCommand(ioStream *iostreams.IOStreams, cmd *cobra.Command, args []string) error {
	manager, err := notification.NewNotificationManager()
	if err != nil {
		return fmt.Errorf("failed to instantiate notification manager: %v", err)
	}
	activeProcesses := manager.GetActiveProcesses()

	// Early return if no active processes.
	if len(activeProcesses) == 0 {
		fmt.Fprintln(ioStream.Out, color.Red("No Running Processes."))
		return nil
	}

	// Show specific notification instance details.
	if len(args) > 0 {
		// Find the running instance.
		instancePid, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("failed to parse argument pid: %v", err)
		}

		instance := manager.GetPid(instancePid)
		if instance == nil {
			fmt.Fprintf(ioStream.Out, color.Red("Instance Pid %d not found."), instancePid)
			return nil
		}
		activeProcesses = []notification.Notification{*instance}
	}

	// List all active notifications.
	nowTime := time.Now()
	for index, instance := range activeProcesses {
		// Time Calculations
		notifyDt := instance.Reminder.EndTime.Sub(nowTime)
		fmt.Fprintf(ioStream.Out, "[%d] ", index)

		// Print Process Structure
		fmt.Fprintf(ioStream.Out, "%s: %s\n", color.Green("Name"), path.Base(instance.PidFilePath))
		fmt.Fprintf(ioStream.Out, "      %s: %d\n", color.Green("Pid"), instance.Pid)
		fmt.Fprintf(ioStream.Out, "      %s:\n", color.Green("Notification"))
		fmt.Fprintf(ioStream.Out, "        %s: %s\n", color.Green("Title"), instance.Notification.Title)
		fmt.Fprintf(ioStream.Out, "        %s: %s\n", color.Green("Summary"), instance.Notification.Summary)
		fmt.Fprintf(ioStream.Out, "        %s: %s\n", color.Green("Icon"), instance.Notification.Icon)
		fmt.Fprintf(ioStream.Out, "        %s: %s\n", color.Green("Urgency"), instance.Notification.Urgency)

		fmt.Fprintf(ioStream.Out, "      %s:\n", color.Green("Reminder"))
		fmt.Fprintf(ioStream.Out, "        %s: %s\n", color.Green("Remind in: "), instance.Reminder.RemindIn.String())
		fmt.Fprintf(ioStream.Out, "        %s: %s\n", color.Green("Time Left: "), notifyDt.String())
		fmt.Fprintf(ioStream.Out, "        %s: %s\n", color.Green("Start Time: "), instance.Reminder.StartTime.Local().Format(time.UnixDate))
		fmt.Fprintf(ioStream.Out, "        %s: %s\n", color.Green("End Time: "), instance.Reminder.EndTime.Local().Format(time.UnixDate))

		fmt.Fprintf(ioStream.Out, "      %s: %s\n", color.Green("Version"), instance.Version)
	}

	return nil
}

func NewListCommand(ioStream *iostreams.IOStreams) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list [PID]",
		Short: "Lists all or given active notification",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleListCommand(ioStream, cmd, args)
		},
	}
	return listCmd
}

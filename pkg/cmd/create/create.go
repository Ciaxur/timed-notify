package create

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/ciaxur/timed-notify/v1/internal/build"
	"github.com/ciaxur/timed-notify/v1/internal/config"
	"github.com/ciaxur/timed-notify/v1/pkg/color"
	"github.com/ciaxur/timed-notify/v1/pkg/iostreams"
	"github.com/ciaxur/timed-notify/v1/pkg/notification"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
)

var (
	title    *string
	summary  *string
	remind   *time.Duration
	icon     *string
	urgency  *int
	isDaemon *bool
)

// Handles Terminating Process Cleanly
func handleInterrupt(pidFile *os.File) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Wait for Interrupt then Clean up
	<-c
	pidFile.Close()
	os.Remove(pidFile.Name())
	os.Exit(0)
}

func handleCreateCommand(ioStream *iostreams.IOStreams, cmd *cobra.Command, args []string) error {
	// Setup temp PID file & directory
	_ = os.Mkdir(config.PidDir, os.ModePerm)
	pidFile, _ := ioutil.TempFile(config.PidDir, "timed-notify.pid.")
	timeNow := time.Now()

	// Create a new notification instance.
	notifyInstance := &notification.Notification{
		PidFilePath: pidFile.Name(),
		Pid:         os.Getpid(),
		Reminder: notification.ReminderStatus{
			RemindIn:  *remind,
			StartTime: timeNow,
			EndTime:   timeNow.Add(*remind),
		},
		Notification: notification.NotificationInfo{
			Title:   *title,
			Summary: *summary,
			Icon:    *icon,
			Urgency: notification.UrgencyIntToString(*urgency),
		},
		Version: build.Version,
	}

	// Save the instance to a pid file.
	notifyInstance.WriteToFile(pidFile)

	// Handle clean up if interrupted.
	// Handle Interrups
	go handleInterrupt(pidFile)

	// Deamonize if Flag
	if *isDaemon {
		fmt.Fprintf(ioStream.Out, color.Magenta("Daemonized Process, running in the Background 😈"))

		// Setup Daemon
		ctx := &daemon.Context{
			WorkDir:     "./",
			Umask:       027,
			Args:        os.Args,
			PidFileName: pidFile.Name(),
			PidFilePerm: os.ModePerm,
		}

		// Check if Debug Mode
		debugStr := cmd.Flags().Lookup("debug").Value.String()
		debugFlag := false
		debugFlag, _ = strconv.ParseBool(debugStr)

		if debugFlag == true {
			ctx.LogFileName = config.BinPath + "/timed-notify.log"
			fmt.Fprintf(ioStream.Out, color.Magenta("Debug logs will be written to: %s"), ctx.LogFileName)
			ctx.LogFilePerm = 0640
		}

		// Release the DAEMON!
		_, err := ctx.Reborn()
		if err != nil {
			return fmt.Errorf("failed to daemonize process: %v", err)
		}
		ctx.Release()
	}

	// Set sleep time to wait prior to invoking notification.
	fmt.Fprintf(ioStream.Out, color.Magenta("🚀 Reminder set for '%s %s' in %.0fs\n"), *summary, *title, remind.Seconds())
	time.Sleep(*remind)

	if err := notifyInstance.TriggerNotification(); err != nil {
		return fmt.Errorf("failed to create notification: %v", err)
	}

	// Clean up.
	pidFile.Close()
	if err := os.Remove(pidFile.Name()); err != nil {
		return fmt.Errorf("failed to remove pid file: %v", err)
	}
	return nil
}

func NewCreateCommand(ioStream *iostreams.IOStreams) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new timed notification",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleCreateCommand(ioStream, cmd, args)
		},
	}

	title = createCmd.PersistentFlags().StringP("title", "t", "Empty Title", "Notification title")
	summary = createCmd.PersistentFlags().StringP("summary", "s", "Emtpy Summary", "Notification summary/body")
	remind = createCmd.PersistentFlags().DurationP("remind", "r", time.Duration(10*time.Second), "Notificaiton reminder.")
	icon = createCmd.PersistentFlags().StringP("icon", "i", config.UserResourcePath+"/Notification.png", "Notification icon.")
	urgency = createCmd.PersistentFlags().IntP("urgency", "u", 2, "Notification urgency (1=low, 2=normal, 3=critical).")
	isDaemon = createCmd.PersistentFlags().BoolP("daemon", "d", false, "Daemonize process or not.")

	return createCmd
}

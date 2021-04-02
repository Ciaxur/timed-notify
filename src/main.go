package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"time"
	"timed-notify/src/Arguments"
	"timed-notify/src/Config"
	"timed-notify/src/Management"

	// External Packages
	"github.com/fatih/color"
	"github.com/sevlyar/go-daemon"
)

// Prints Help Menu
func printHelp() {
	cyan := color.New(color.FgHiCyan).SprintFunc()

	Config.InfoOut.Print("Usage timed-notify:\n\t")
	fmt.Printf(cyan("timed-notify") + " <SUMMARY> [OPTIONS] - Create a Timed Notification\n")

	Config.InfoOut.Printf("Help Options:\n")
	fmt.Printf("\t-h, -Help \t\t\t Displays Help Menu\n")
	fmt.Printf("\t-v, -Version \t\t\t Displays Version\n")
	fmt.Printf("\t-Debug \t\t\t\t Enables Log Output\n")

	Config.InfoOut.Printf("\nNotification Options:\n")
	fmt.Printf("\t-t, -Title \t\t\t Sets the Notification Title\n")
	fmt.Printf("\t-m, -Summary \t\t\t Sets the Notification Summary\n")
	fmt.Printf("\t-r, -Remind \t\t\t Sets the Notification Delay Time, Default=10s\n")
	fmt.Printf("\t-i, -Icon \t\t\t Sets the Notification Icon, Default [" + Config.BinPath + "/Notification.png]\n")
	fmt.Printf("\t-u, -Urgency \t\t\t Sets the Notification Urgency [1=Low, 2=Normal, 3=Critical]\n")
	fmt.Printf("\t-d, -Daemon \t\t\t Runs Process as a Daemon\n")

	Config.InfoOut.Printf("\nManagement Options:\n")
	fmt.Printf("\t-l, -List \t\t\t Lists Running/Pending Processes/Daemons\n")
	fmt.Printf("\t-p, -PidPrint \t\t\t Lists Given Pid Information\n")
	fmt.Printf("\t-k, -Kill \t\t\t Terminates Given Pid\n")

	Config.InfoOut.Println("\nExamples: ")
	Config.StdOut.Printf("\ttimed-notify %s \n", cyan("-Remind {[Interger][s/m/h]} -Title {message}"))
	Config.StdOut.Printf("\ttimed-notify %s \n", cyan("-Remind {[Interger][s/m/h]} -Title {message} -Summary {summary}"))
	Config.StdOut.Printf("\ttimed-notify %s \n", cyan("-r {[Interger][s/m/h]} -t {title} -m {summary}"))
	Config.StdOut.Printf("\ttimed-notify %s \n", cyan("-r {[Interger][s/m/h]} -t {title} -m {summary} -u 3"))
}

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

func main() {
	// Parse Arguments
	args := Arguments.ParseInput()

	// VERIFY ARGUMENTS
	if args.IsHelp { // Print Help Menu
		printHelp()
		os.Exit(0)
	} else if args.IsVersion { // Print Program Version
		fmt.Printf("timed-notify Version %s\n", Config.VERSION)
		os.Exit(0)
	}

	// Redirect to Management
	Management.HandleManagementArgs(args)

	// Title and Reminder must be enabled
	if args.Title == "" && args.Summary == "" {
		Config.ErrOut.Print("Title or Summary of notification MUST be set!\n\n")
		printHelp()
		os.Exit(1)
	}

	// DETERMINE URGENCY
	var urgentLevel = "normal"
	switch args.Urgency {
	case 1:
		urgentLevel = "low"
	case 2:
		urgentLevel = "normal"
	case 3:
		urgentLevel = "critical"
	}

	// Setup Temp PID File & Directory
	_ = os.Mkdir(Config.PidDir, os.ModePerm)
	pidFile, _ := ioutil.TempFile(Config.PidDir, "timed-notify.pid.")
	timeNow := time.Now()

	// Store Data in Pid File
	Management.PidFileWrite(&Management.ProcessStatus{
		PidFilePath: pidFile.Name(),
		Pid:         os.Getpid(),
		Reminder: Management.ReminderStatus{
			RemindIn:  args.Remind,
			StartTime: timeNow,
			EndTime:   timeNow.Add(args.Remind),
		},
		Notification: Management.NotificationInfo{
			Title:   args.Title,
			Summary: args.Summary,
			Icon:    args.Icon,
			Urgency: urgentLevel,
		},
		Version: Config.VERSION,
	}, pidFile)

	// Handle Interrups
	go handleInterrupt(pidFile)

	// Deamonize if Flag
	if args.IsDaemon {
		Config.InfoOut.Println("Daemonized Process, running in the Background 😈")

		// Setup Daemon
		ctx := &daemon.Context{
			WorkDir:     "./",
			Umask:       027,
			Args:        os.Args,
			PidFileName: pidFile.Name(),
			PidFilePerm: os.ModePerm,
		}

		// Check if Debug Mode
		if args.IsLog {
			Config.InfoOut.Println("Debug Logs will be saved in: " + Config.BinPath)
			ctx.LogFileName = Config.BinPath + "/timed-notify.log"
			ctx.LogFilePerm = 0640
		}

		// Release the DAEMON!
		d, err := ctx.Reborn()
		if err != nil {
			Config.ErrOut.Printf("Unable to run: %s\n", err)
		}
		if d != nil {
			os.Exit(0)
		}
		ctx.Release()
	}

	// SET SLEEP TIME
	Config.InfoOut.Printf("🚀 Reminder set for '%s %s' in %.0fs\n", args.Summary, args.Title, args.Remind.Seconds())
	time.Sleep(args.Remind)

	// INITIATE NOTIFICATION
	cmd := exec.Command("notify-send", args.Title, args.Summary, "-i", args.Icon, "-u", urgentLevel)
	cmd.Start()

	// Clean up
	pidFile.Close()
	os.Remove(pidFile.Name())
}

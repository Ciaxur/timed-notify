package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"timed-notify/src/Arguments"
	"timed-notify/src/Config"

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
	fmt.Printf("\t-Debug \t\t\t Enables Log Output\n")

	Config.InfoOut.Printf("\nNotification Options:\n")
	fmt.Printf("\t-t, -Title \t\t\t Sets the Notification Title\n")
	fmt.Printf("\t-m, -Summary \t\t\t Sets the Notification Summary\n")
	fmt.Printf("\t-r, -Remind \t\t\t Sets the Notification Delay Time, Default=10s\n")
	fmt.Printf("\t-i, -Icon \t\t\t Sets the Notification Icon, Default [" + Config.BinPath + "/Notification.png]\n")
	fmt.Printf("\t-u, -Urgency \t\t\t Sets the Notification Urgency [1=Low, 2=Normal, 3=Critical]\n")
	fmt.Printf("\t-d, -Daemon \t\t\t Runs Process as a Daemon\n")

	Config.InfoOut.Println("\nExamples: ")
	Config.StdOut.Printf("\ttimed-notify %s \n", cyan("-Remind {[Interger][s/m/h]} -Title {message}"))
	Config.StdOut.Printf("\ttimed-notify %s \n", cyan("-Remind {[Interger][s/m/h]} -Title {message} -Summary {summary}"))
	Config.StdOut.Printf("\ttimed-notify %s \n", cyan("-r {[Interger][s/m/h]} -t {title} -m {summary}"))
	Config.StdOut.Printf("\ttimed-notify %s \n", cyan("-r {[Interger][s/m/h]} -t {title} -m {summary} -u 3"))
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

	// Deamonize if Flag
	if args.IsDaemon {
		Config.InfoOut.Println("Daemonized Process, running in the Background 😈")

		// Setup Daemon
		ctx := &daemon.Context{
			WorkDir: "./",
			Umask:   027,
			Args:    os.Args,
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
}

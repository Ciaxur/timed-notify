package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	// External Packages
	"github.com/fatih/color"
	"github.com/sevlyar/go-daemon"
)

// CONFIGURE GLOBAL STD OUTPUT COLORS
var (
	errOut     = color.New(color.FgRed).Add(color.Bold)
	infoOut    = color.New(color.FgHiMagenta)
	stdOut     = color.New()
	binPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	VERSION    = "1.0.2"
)

// Structure for Valid CLI Arguments
type cliArguments struct {
	remind    time.Duration
	title     string
	summary   string
	icon      string
	urgency   int
	isDaemon  bool
	isVersion bool
	isHelp    bool
}

// Parses throught the CLI Arguments
// Returning necessary Flags
func parseInput() cliArguments {
	var FlagRemind = flag.Duration("Remind", time.Duration(math.Pow10(10)), "Time to Remind")
	flag.DurationVar(FlagRemind, "r", time.Duration(math.Pow10(10)), "Timer to Remind")

	var FlagTitle = flag.String("Title", "Timed Notify", "Message for title")
	flag.StringVar(FlagTitle, "t", "Timed Notify", "Message for Title")

	var FlagSummary = flag.String("Summary", "<no body>", "Message for summary")
	flag.StringVar(FlagSummary, "m", "<no body>", "Message for summary")

	var FlagIcon = flag.String("Icon", "/usr/share/timed-notify/res/Notification.png", "Custom Icon to use")
	flag.StringVar(FlagIcon, "i", "/usr/share/timed-notify/res/Notification.png", "Custom Icon to use")

	var FlagUrgent = flag.Int("Urgency", 2, "Set urgancy level")
	flag.IntVar(FlagUrgent, "u", 2, "Set urgancy level")

	var FlagBool = flag.Bool("Daemon", false, "Daemonize process or not")
	flag.BoolVar(FlagBool, "d", false, "Daemonize process or not")

	var FlagVersion = flag.Bool("Version", false, "Displays the current timed-notify Version")
	flag.BoolVar(FlagVersion, "v", false, "Displays the current timed-notify Version")

	var FlagHelp = flag.Bool("Help", false, "Displays Help Menu")
	flag.BoolVar(FlagHelp, "h", false, "Displays Help Menu")

	flag.Parse()

	flags := cliArguments{*FlagRemind, *FlagTitle, *FlagSummary, *FlagIcon, *FlagUrgent, *FlagBool, *FlagVersion, *FlagHelp}
	return flags
}

// Prints Help Menu
func printHelp() {
	cyan := color.New(color.FgHiCyan).SprintFunc()

	infoOut.Print("Usage timed-notify:\n\t")
	fmt.Printf(cyan("timed-notify") + " <SUMMARY> [OPTIONS] - Create a Timed Notification\n")

	infoOut.Printf("Help Options:\n")
	fmt.Printf("\t-h, -Help \t\t\t Displays Help Menu\n")
	fmt.Printf("\t-v, -Version \t\t\t Displays Version\n")

	infoOut.Printf("\nNotification Options:\n")
	fmt.Printf("\t-t, -Title \t\t\t Sets the Notification Title\n")
	fmt.Printf("\t-m, -Summary \t\t\t Sets the Notification Summary\n")
	fmt.Printf("\t-r, -Remind \t\t\t Sets the Notification Delay Time, Default=10s\n")
	fmt.Printf("\t-i, -Icon \t\t\t Sets the Notification Icon, Default [" + binPath + "/Notification.png]\n")
	fmt.Printf("\t-u, -Urgency \t\t\t Sets the Notification Urgency [1=Low, 2=Normal, 3=Critical]\n")
	fmt.Printf("\t-d, -Daemon \t\t\t Runs Process as a Daemon\n")

	infoOut.Println("\nExamples: ")
	stdOut.Printf("\ttimed-notify %s \n", cyan("-Remind {[Interger][s/m/h]} -Title {message}"))
	stdOut.Printf("\ttimed-notify %s \n", cyan("-Remind {[Interger][s/m/h]} -Title {message} -Summary {summary}"))
	stdOut.Printf("\ttimed-notify %s \n", cyan("-r {[Interger][s/m/h]} -t {title} -m {summary}"))
	stdOut.Printf("\ttimed-notify %s \n", cyan("-r {[Interger][s/m/h]} -t {title} -m {summary} -u 3"))
}

// Simple wrapper that returns the conversion of string to int
func getIntStr(sVal string) int {
	intVal, err := strconv.Atoi(sVal)
	if err != nil {
		errOut.Println("First Argument is time to Sleep! [int]")
		os.Exit(1)
	}
	return intVal
}

func main() {
	// Parse Arguments
	args := parseInput()

	// VERIFY ARGUMENTS
	if args.isHelp { // Print Help Menu
		printHelp()
		os.Exit(0)
	} else if args.isVersion { // Print Program Version
		fmt.Printf("timed-notify Version %s\n", VERSION)
		os.Exit(0)
	}

	// Title and Reminder must be enabled
	if args.title == "" && args.summary == "" {
		errOut.Print("Title or Summary of notification MUST be set!\n\n")
		printHelp()
		os.Exit(1)
	}

	// DETERMINE URGENCY
	var urgentLevel = "normal"
	switch args.urgency {
	case 1:
		urgentLevel = "low"
	case 2:
		urgentLevel = "normal"
	case 3:
		urgentLevel = "critical"

	}

	// Deamonize if Flag
	if args.isDaemon {
		infoOut.Println("Daemonized Process, running in the Background 😈")

		// Setup Daemon
		ctx := &daemon.Context{
			LogFileName: binPath + "/timed-notify.log",
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       027,
			Args:        os.Args,
		}

		// Release the DAEMON!
		d, err := ctx.Reborn()
		if err != nil {
			errOut.Printf("Unable to run: %s\n", err)
		}
		if d != nil {
			os.Exit(0)
		}
		ctx.Release()
	}

	// SET SLEEP TIME
	infoOut.Printf("🚀 Reminder set for '%s %s' in %.0fs\n", args.summary, args.title, args.remind.Seconds())
	time.Sleep(args.remind)

	// INITIATE NOTIFICATION
	cmd := exec.Command("notify-send", args.title, args.summary, "-i", args.icon, "-u", urgentLevel)
	cmd.Start()
}

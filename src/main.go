package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	// "strings"
	"time"
	"flag"
	"fmt"

	// External Packages
	"github.com/sevlyar/go-daemon"
	"github.com/fatih/color"
)


// CONFIGURE GLOBAL STD OUTPUT COLORS
var (
	errOut  = color.New(color.FgRed).Add(color.Bold)
	infoOut = color.New(color.FgHiMagenta)
	stdOut  = color.New()
	// Obtain Icon Full Path
	 binPath,_ = filepath.Abs(filepath.Dir(os.Args[0]))
)


type commandline_arguments struct{
	 remind string
	 title string
	 summary string
	 icon string
	 urgency int
	 daemon bool
}
func parseInput() commandline_arguments{
	var FlagRemind=flag.String("Remind", "", "Time to Remind")
	flag.StringVar(FlagRemind,"r", "", "Timer to Remind")

	var FlagTitle=flag.String("Title", "", "Message for title")
	flag.StringVar(FlagTitle,"t", "", "Message for Title")

	var FlagSummary=flag.String("Summary", "", "Message for summary")
	flag.StringVar(FlagSummary,"s", "", "Message for summary")

	var FlagIcon=flag.String("Icon", binPath+"/Notification.png", "Custom Icon to use")
	flag.StringVar(FlagIcon,"i", binPath+"/Notification.png", "Custom Icon to use")

	var FlagUrgent=flag.Int("Urgency", 2, "Set urgancy level")
	flag.IntVar(FlagUrgent, "u", 2, "Set urgancy level")

	var FlagBool=flag.Bool("Daemon", false, "Daemonize process or not")
	flag.BoolVar(FlagBool, "d", false, "Daemonize process or not")

	flag.Parse()
	flags := commandline_arguments{*FlagRemind, *FlagTitle,*FlagSummary,*FlagIcon,*FlagUrgent,*FlagBool}
	return flags
}
// Prints Help Menu
func printHelp() {
	cyan := color.New(color.FgHiCyan).SprintFunc()

	errOut.Println("Two Arguments Required:")
	infoOut.Println("\t-Remind: [Time {amount(s/m/h)}]")
	infoOut.Println("\t-Title: [Message]")

	infoOut.Println("Examples: ")
	stdOut.Printf("\tapp %s \n", cyan("-Remind {[Interger][s/m/h]} -Title {message}"))
	stdOut.Printf("\tapp %s \n", cyan("-Remind 2s -Title \"Hello World\" -Summary \"Here we go!\""))
	stdOut.Printf("\tapp %s \n", cyan("-Title \"Hello World\" -Remind 2s -Urgency 3"))
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
	var args = parseInput()
	iconPath := args.icon
	Remind := args.remind
	Summary := args.summary
	Title := args.title
	Urgency := args.urgency
	isDaemon := args.daemon

	// VERIFY ARGUMENTS 
	// Title and Reminder must be enabled
	if(Title ==""){
		printHelp()
		errOut.Println("Title of notification is not set")
		os.Exit(-1)
	} else if (Remind == "") {
		printHelp()
		errOut.Println("Reminder time is not set")
		os.Exit(-1)
	}
	// DETERMINE SLEEP AMOUNT
	var dTime time.Duration
	tTypeStr := Remind[len(Remind)-1]
	waitTime := Remind[:len(Remind)-1]
	waitType := "Seconds"

	switch tTypeStr {
	case 's': // Specifically Seconds
		dTime = time.Duration(getIntStr(waitTime)) * time.Second
	case 'm': // Minutes
		dTime = time.Duration(getIntStr(waitTime)) * time.Minute
		waitType = "Minutes"
	case 'h': // Hours
		dTime = time.Duration(getIntStr(waitTime)) * time.Hour
		waitType = "Hours"
	default: // Defaulted to Seconds
		waitTime = os.Args[1]
		dTime = time.Duration(getIntStr(waitTime)) * time.Second
	}
	var urgentLevel="normal"
	switch Urgency{
	case 1:
		urgentLevel="low"
	case 2:
		urgentLevel="normal"
	case 3:
		urgentLevel="critical"

	}
	infoOut.Printf("Waiting for %s %s to output '%s'\n", waitTime, waitType, os.Args[2])


	// Deamonize if Flag
	if isDaemon {
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
	time.Sleep(dTime)

	// INITIATE NOTIFICATION
	fmt.Println("Sending notify-send", Title, Summary, "-i", iconPath, "-u", urgentLevel)
	cmd := exec.Command("notify-send", Title, Summary, "-i", iconPath, "-u", urgentLevel)
	cmd.Start()
}

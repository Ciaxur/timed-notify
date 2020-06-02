package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sevlyar/go-daemon"

	"github.com/fatih/color"
)

// CONFIGURE GLOBAL STD OUTPUT COLORS
var (
	errOut  = color.New(color.FgRed).Add(color.Bold)
	infoOut = color.New(color.FgHiMagenta)
	stdOut  = color.New()
)

// Prints Help Menu
func printHelp() {
	cyan := color.New(color.FgHiCyan).SprintFunc()

	errOut.Println("Two Arguments Required:")
	infoOut.Println("\tArgument 1: [Time {amount(s/m/h)}]")
	infoOut.Println("\tArgument 2: [Message]")

	infoOut.Println("Examples: ")
	stdOut.Printf("\tapp %s \n", cyan("{time[s/m/h]} {message}"))
	stdOut.Printf("\tapp %s \n", cyan("2s \"Hello World\""))
	stdOut.Printf("\tapp %s \n", cyan("2 \"Hello World\""))
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
	// Keep track of Arguments
	isDaemon := false
	args := []string{}

	// SEPERATE ARGUMENT FLAGS & REGULAR ARGUMENTS
	for _, elt := range os.Args {
		// Check for Flags
		if len(elt) == 2 && elt[0] == '-' { // Make sure argument is a flag
			if elt == "-d" { // Daemon Flag
				isDaemon = true
			} else { // Unknown Flag
				errOut.Print("Unknown Argument Flag: ")
				infoOut.Printf("%s\n", elt)
			}
		} else {
			args = append(args, elt)
		}
	}
	os.Args = args // Assign Regular Arugments

	// VERIFY ARGUMENTS (3 Arguments : Prog, Seconds, Message)
	if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "help" {
		printHelp()
		os.Exit(0)
	} else if len(os.Args) != 3 {
		errOut.Println("Not Enough Arguments!")
		printHelp()
		os.Exit(1)
	}

	// DETERMINE SLEEP AMOUNT
	var dTime time.Duration
	tTypeStr := os.Args[1][len(os.Args[1])-1]
	waitTime := os.Args[1][:len(os.Args[1])-1]
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
	infoOut.Printf("Waiting for %s %s to output '%s'\n", waitTime, waitType, os.Args[2])

	// Obtain Icon Full Path
	binPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	iconPath := binPath + "/Notification.png"

	// Deamonize if Flag
	if isDaemon {
		infoOut.Println("Daemonized Process, running in the Background 😈")

		// Setup Daemon
		ctx := &daemon.Context{
			PidFileName: binPath + "/timed-notify.pid",
			PidFilePerm: 0644,
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
	cmd := exec.Command("notify-send", os.Args[2], "-i", iconPath)
	cmd.Start()
}

package Management

import (
	"log"
	"os"
	"path"
	"time"
	"timed-notify/src/Arguments"
	"timed-notify/src/Config"

	"github.com/fatih/color"
)

// Helper Function to Panic on Error
func handlePanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Prints given Process Status Structure information
//  @param nowTime - The Current Time
//  @param pidFileName - Name of the PID File
//  @param procStruct - ProcessStatus Structure for PID
//  @param outColor - The Color to Output Information about PID Status
func printPidInfo(nowTime time.Time, pidFileName string, procStruct *ProcessStatus, outColor *color.Color) {
	// Time Calculations
	notifyDt := procStruct.Reminder.EndTime.Sub(nowTime)

	// Print Process Structure
	Config.StdOut.Printf("%s: %s\n", outColor.Sprintf("Name"), pidFileName)
	Config.StdOut.Printf("      %s: %d\n", outColor.Sprintf("Pid"), procStruct.Pid)

	Config.StdOut.Printf("      %s:\n", outColor.Sprintf("Notification"))
	Config.StdOut.Printf("        %s: %s\n", outColor.Sprintf("Title"), procStruct.Notification.Title)
	Config.StdOut.Printf("        %s: %s\n", outColor.Sprintf("Summary"), procStruct.Notification.Summary)
	Config.StdOut.Printf("        %s: %s\n", outColor.Sprintf("Icon"), procStruct.Notification.Icon)
	Config.StdOut.Printf("        %s: %s\n", outColor.Sprintf("Urgency"), procStruct.Notification.Urgency)

	Config.StdOut.Printf("      %s:\n", outColor.Sprintf("Reminder"))
	Config.StdOut.Printf("        %s: %s\n", outColor.Sprintf("Remind in: "), procStruct.Reminder.RemindIn.String())
	Config.StdOut.Printf("        %s: %s\n", outColor.Sprintf("Time Left: "), notifyDt.String())
	Config.StdOut.Printf("        %s: %s\n", outColor.Sprintf("Start Time: "), procStruct.Reminder.StartTime.Local().Format(time.UnixDate))
	Config.StdOut.Printf("        %s: %s\n", outColor.Sprintf("End Time: "), procStruct.Reminder.EndTime.Local().Format(time.UnixDate))

	Config.StdOut.Printf("      %s: %s\n", outColor.Sprintf("Version"), procStruct.Version)
}

// Helper function to check if PID Directory is available
//  Outputs to stdout if no PID Directory found, which means
//  that no PIDs are Running
func checkPidDirectory() {
	// Verify Directory Exists
	if _, err := os.Stat(Config.PidDir); os.IsNotExist(err) {
		Config.ErrOut.Printf("No Running Processes\n")
		return
	}
}

// Lists all running PID timed-notify Processes' Information
func listRunningPid() {
	// Verify Existing Directory
	checkPidDirectory()

	// Read Pid Files
	pidFiles, _ := os.ReadDir(Config.PidDir)

	// Output Settings
	green := color.New(color.FgGreen)

	// Output Summary
	Config.InfoOut.Printf("Listing Running Pids:\n")
	Config.StdOut.Printf("  %s: %d\n", green.Sprintf("Running Processes"), len(pidFiles))
	nowTime := time.Now()

	for index, pid := range pidFiles {
		fullPath := path.Join(Config.PidDir, pid.Name())
		procStruct, err := PidFileRead(fullPath)
		handlePanic(err)

		// Output Process Info
		Config.StdOut.Printf("[%d] ", index)
		printPidInfo(nowTime, pid.Name(), procStruct, green)
	}
}

// Tries to find PID for timed-notify, returning it's
//  current Status
// Outputs to stdout if PID not found
// @param pid - Process ID for timed-notify running process
// @returns
//  - ProcessStatus of given PID; nil if not found
//  - File Name for PID File
func getPidProcessInfo(pid int) (*ProcessStatus, string) {
	// Verify Existing Directory
	checkPidDirectory()

	// Read Pid Files
	pidFiles, _ := os.ReadDir(Config.PidDir)

	// Find & Print Pid Info
	for _, pidFile := range pidFiles {
		fullPath := path.Join(Config.PidDir, pidFile.Name())
		procStruct, err := PidFileRead(fullPath)
		handlePanic(err)

		if procStruct.Pid == pid {
			return procStruct, pidFile.Name()
		}
	}

	Config.ErrOut.Printf("PID %d is not running.\n", pid)
	return nil, ""
}

// Lists given PID's Information
//  @param pid - Process ID for a running timed-notify Process
func listPidInfo(pid int) {
	if pidInfo, pidFileName := getPidProcessInfo(pid); pidInfo != nil {
		Config.StdOut.Printf("    ")
		printPidInfo(time.Now(), pidFileName, pidInfo, color.New(color.FgGreen))
	}
}

// HandleManagementArgs -
//  Handles passed in Arguments Accordingly
//  @param args - Parsed CLI Arguments
func HandleManagementArgs(args Arguments.CliArguments) {
	// List Running PIDS
	if args.IsListPids {
		listRunningPid()
		os.Exit(0)
	} else if args.PrintPid > -1 {
		listPidInfo(args.PrintPid)
		os.Exit(0)
	} else if args.KillPid > -1 {
		// Make sure Process is part of timed-notify
		if pidInfo, _ := getPidProcessInfo(args.KillPid); pidInfo != nil {
			// Find and Issue an Interrupt for Process
			proc, err := os.FindProcess(pidInfo.Pid)
			if err != nil {
				Config.ErrOut.Printf("PID %d could not be found.\n", pidInfo.Pid)
			} else {
				Config.InfoOut.Printf("PID %d Interrupted.\n", pidInfo.Pid)
				proc.Signal(os.Interrupt)
			}
		}
		os.Exit(0)
	}
}

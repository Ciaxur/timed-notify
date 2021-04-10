package Arguments

import (
	"flag"
	"math"
	"runtime"
	"time"
	"timed-notify/src/Config"
)

// CliArguments -
//  Structure for Valid CLI Arguments
type CliArguments struct {
	Remind    time.Duration
	Title     string
	Summary   string
	Icon      string
	Urgency   int
	IsDaemon  bool
	IsLog     bool
	IsVersion bool
	IsHelp    bool

	// Management Flags
	IsListPids bool
	PrintPid   int
	KillPid    int
}

// ParseArguments -
//  Parses throught the CLI Arguments
//  Returning necessary Flags
func ParseInput() CliArguments {
	// NOTIFICATION OPTIONS
	var FlagRemind = flag.Duration("Remind", time.Duration(math.Pow10(10)), "Time to Remind")
	flag.DurationVar(FlagRemind, "r", time.Duration(math.Pow10(10)), "Timer to Remind")

	var FlagTitle = flag.String("Title", "Timed Notify", "Message for title")
	flag.StringVar(FlagTitle, "t", "Timed Notify", "Message for Title")

	var FlagSummary = flag.String("Summary", "<no body>", "Message for summary")
	flag.StringVar(FlagSummary, "m", "<no body>", "Message for summary")

	// Default Icon Location
	if runtime.GOOS == "darwin" { // MacOS Default Shared Path
		Config.ResPath = "/usr/local/share/timed-notify"
	}

	var FlagIcon = flag.String("Icon", Config.ResPath+"/Notification.png", "Custom Icon to use")
	flag.StringVar(FlagIcon, "i", Config.ResPath+"/Notification.png", "Custom Icon to use")

	var FlagUrgent = flag.Int("Urgency", 2, "Set urgancy level")
	flag.IntVar(FlagUrgent, "u", 2, "Set urgancy level")

	var FlagDaemon = flag.Bool("Daemon", false, "Daemonize process or not")
	flag.BoolVar(FlagDaemon, "d", false, "Daemonize process or not")

	var FlagLog = flag.Bool("Debug", false, "Enables Daemonize log output")

	var FlagVersion = flag.Bool("Version", false, "Displays the current timed-notify Version")
	flag.BoolVar(FlagVersion, "v", false, "Displays the current timed-notify Version")

	var FlagHelp = flag.Bool("Help", false, "Displays Help Menu")
	flag.BoolVar(FlagHelp, "h", false, "Displays Help Menu")

	// MANAGEMENT OPTIONS
	var FlagListPids = flag.Bool("List", false, "Lists Running/Pending Processes/Daemons")
	flag.BoolVar(FlagListPids, "l", false, "Lists Running/Pending Processes/Daemons")

	var FlagPrintPid = flag.Int("PidPrint", -1, "Lists Given Pid Information")
	flag.IntVar(FlagPrintPid, "p", -1, "Lists Given Pid Information")

	var FlagKillPid = flag.Int("Kill", -1, "Terminates Given Pid")
	flag.IntVar(FlagKillPid, "k", -1, "Terminates Given Pid")

	// PARSE & APPLY
	flag.Parse()
	flags := CliArguments{
		*FlagRemind,
		*FlagTitle,
		*FlagSummary,
		*FlagIcon,
		*FlagUrgent,
		*FlagDaemon,
		*FlagLog,
		*FlagVersion,
		*FlagHelp,
		*FlagListPids,
		*FlagPrintPid,
		*FlagKillPid,
	}
	return flags
}

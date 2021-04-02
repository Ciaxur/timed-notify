package Arguments

import (
	"flag"
	"math"
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
}

// ParseArguments -
//  Parses throught the CLI Arguments
//  Returning necessary Flags
func ParseInput() CliArguments {
	var FlagRemind = flag.Duration("Remind", time.Duration(math.Pow10(10)), "Time to Remind")
	flag.DurationVar(FlagRemind, "r", time.Duration(math.Pow10(10)), "Timer to Remind")

	var FlagTitle = flag.String("Title", "Timed Notify", "Message for title")
	flag.StringVar(FlagTitle, "t", "Timed Notify", "Message for Title")

	var FlagSummary = flag.String("Summary", "<no body>", "Message for summary")
	flag.StringVar(FlagSummary, "m", "<no body>", "Message for summary")

	var FlagIcon = flag.String("Icon", Config.ResPath+"/res/Notification.png", "Custom Icon to use")
	flag.StringVar(FlagIcon, "i", Config.ResPath+"/res/Notification.png", "Custom Icon to use")

	var FlagUrgent = flag.Int("Urgency", 2, "Set urgancy level")
	flag.IntVar(FlagUrgent, "u", 2, "Set urgancy level")

	var FlagBool = flag.Bool("Daemon", false, "Daemonize process or not")
	flag.BoolVar(FlagBool, "d", false, "Daemonize process or not")

	var FlagLog = flag.Bool("Debug", false, "Enables Daemonize log output")

	var FlagVersion = flag.Bool("Version", false, "Displays the current timed-notify Version")
	flag.BoolVar(FlagVersion, "v", false, "Displays the current timed-notify Version")

	var FlagHelp = flag.Bool("Help", false, "Displays Help Menu")
	flag.BoolVar(FlagHelp, "h", false, "Displays Help Menu")

	flag.Parse()

	flags := CliArguments{*FlagRemind, *FlagTitle, *FlagSummary, *FlagIcon, *FlagUrgent, *FlagBool, *FlagLog, *FlagVersion, *FlagHelp}
	return flags
}

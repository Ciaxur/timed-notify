# Timed-Notify
Timed Notifications for Linux Notification Daemons

## Dependencies 📦
**`timed-notify`** uses Linux's `notify-send` program to send notification information to the systems's Notification Daemon.

**Golang Packages used**
- [sevylar's go-daemon](https://github.com/sevlyar/go-daemon) - Daemonize Processes
- [fatih's color](https://github.com/fatih/color) - Colorized Prints

## Build 🔨
Building the packages using Go
You can place any `Notification Icon` into the base directory of the binary to enable icon for the notification

```bash
go build -o ./bin/app ./src         # Builds the Entire Program into bin/app
cp ./res/Notification.png ./bin/    # Copy Default Notification Icon into base directory of app
```

## Running 🚀
Running just the base `./bin/app` will print out the help menu and examples

```bash
# These Wait in the Terminal
./bin/app -Remind 2s  -Title "Hello World"      # Notifies "Hello World" in 2 Seconds
./bin/app -r 10m -t "Hello World" -summary "My name is Omar^2"      # Notifies "Hello World" in 10 Minutes
./bin/app -t "Hello World" -r 10h      # Notifies "Hello World" in 1 Hour

# Daemonizing the Program (Run in Background)
/bin/app -t "Hello World" -r 1s -d

```

**Expected Parameters**
- `-r/-Remind` - Time to wait to notify, identified as *amount*(s/m/h)
  - `amount` = Integer Value
  - `s` = Seconds
  - `m` = Minutes
  - `h` = Hours
- `-t/-Title` - The String Message in the Notification
- `-s/-Summary` - The String Summary of Notifcation
- `-u/-Urgency` - Urgency level of Notification
- `-i/-Icon` - Image to use as icon for the Notification
- `-t/-Title` - The String Message in the Notification

**Argument Flags**
- `-d/--Daemon` - Daemonizes the Process to run in the background

## Logs and Pid 🧾
Log File is saved under `bin/timed-notify.log`

## License 📝
This project is licensed under [MIT](LICENSE).

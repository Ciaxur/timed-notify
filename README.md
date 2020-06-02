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
./bin/app 2s  "Hello World"      # Notifies "Hello World" in 2 Seconds
./bin/app 10m "Hello World"      # Notifies "Hello World" in 10 Minutes
./bin/app 1h  "Hello World"      # Notifies "Hello World" in 1 Hour

# Daemonizing the Program (Run in Background)
./bin/app 2s  "Hello World" -d
```

**Expected Parameters**
- `Param1` - Time to wait to notify, identified as *amount*(s/m/h)
  - `amount` = Integer Value
  - `s` = Seconds
  - `m` = Minutes
  - `h` = Hours
- `Param2` - The String Message in the Notification

**Argument Flags**
- `-d` - Daemonizes the Process to run in the background

## Logs and Pid 🧾
Log File is saved under `bin/timed-notify.log`

## License 📝
This project is licensed under [MIT](LICENSE).

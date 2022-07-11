<h1 align="center">
Timed-Notify
</h1>
<p align="center">
  <img src="res/Bell.png" height=32>
</p>
<p align="center">
 Timed Notifications for <b>Mac & Linux</b> Notification Daemons
</p>

## 📦 Dependencies
**`timed-notify`** uses Linux's `notify-send` program to send notification
information to the systems's Notification Daemon.

## 💎 Installing
In the `scripts` Directory, there is an [install](scripts/install.sh) and
[uninstall](scripts/uninstall.sh) script to help ease the process

```bash
cd ./scripts && ./install.sh     # Builds and Installs Binary and Resources
cd ./scripts && ./uninstall.sh   # Removes Binary and Resources
```


## 🔨 Build
Building the packages using Go
You can place any `Notification Icon` into the base directory of the binary to
enable icon for the notification.

Using [build.sh](scripts/build.sh) will compile the binary into the "build"
directory.
```bash
./scripts/build.sh
```

## 🚀 Running
Running just the base `timed-notify` will print out the help menu and examples

```bash
# These Wait in the Terminal
# Notifies "Hello World" in 2 Seconds
timed-notify create --remind 2s  --title "Hello World
# Notifies "Hello World" in 10 Minutes
timed-notify create -r 10m -t "Hello World" -s "My name is Omar^2
# Notifies "Hello World" in 1 Hour
timed-notify create -t "Hello World" -r 1h

# Daemonizing the Program (Run in Background)
timed-notify create -t "Hello World" -r 1s -d
```

Help menu.

```bash
Issue and manage timed desktop notifications

Usage:
  timed-notify [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create a new timed notification
  help        Help about any command
  kill        Terminates the given Pid
  list        Lists all or given active notification
  version     App version

Flags:
      --debug   Enables daemonized log output to /tmp.
  -h, --help    help for timed-notify

Use "timed-notify [command] --help" for more information about a command.
```

## 🧾 Logs and Pid
Log File is saved under `[bin]/timed-notify.log`
  - `bin`: Where the binary is stored

## 📝 License
This project is licensed under [MIT](LICENSE).

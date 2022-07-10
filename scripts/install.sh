#/usr/bin/env bash

# Project paths.
ROOT_DIR=$(realpath `dirname $0`/..)
RESOURCES_DIR="$ROOT_DIR/res"
SCRIPTS_DIR="$ROOT_DIR/scripts"

# Default Install Path (LINUX).
BIN_PATH="/usr/bin"
RES='/usr/share/timed-notify'

# Install Path for MacOS
if [ `uname | grep Darwin` == "Darwin" ]; then
  echo "MacOS Install"
  BIN_PATH="/usr/local/bin"
  RES='/usr/local/share/timed-notify'
fi

# Populate install share directory (resources).
echo "Creating resources directory '$RES'"
sudo mkdir -p $RES

# Build binary and install on system.
echo "Installing app '$BIN_PATH/timed-notify'."
bash $SCRIPTS_DIR/build.sh
if [ $? != 0 ]; then echo "Failed to build go package!" && exit 1; fi

sudo mv "$RESOURCES_DIR/timed-notify" "$BIN_PATH/timed-notify"
if [ $? != 0 ]; then echo "Failed to install app in install path!" && exit 1; fi

# Copy Resources.
echo "Copying Resources to '$RES.'"
sudo cp "$RESOURCES_DIR/Notification.png" "$RES/"
if [ $? != 0 ]; then echo "Failed to copy resources!" && exit 1; fi

echo "timed-notify Successfully installed! 🚀"
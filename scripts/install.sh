#/bin/sh

# Default Path (LINUX)
PATH_TO_RES="$(dirname $0)/../res"
BIN_PATH="/usr/bin"
RES='/usr/share/timed-notify'

# MacOS
if [ `uname | grep Darwin` == "Darwin" ]; then
  echo "MacOS Install"
  BIN_PATH="/usr/local/bin"
  RES='/usr/local/share/timed-notify'
fi

# Populate Directory
echo "Creating Resources Directory '$RES'"
sudo mkdir -p $RES

# Build Package
echo "Building Package to '$BIN_PATH/timed-notify'..."
go mod tidy || go get -d ./src
go build -o "$PATH_TO_RES/timed-notify" "./src" #|| 
if [ $? != 0 ]; then echo "Failed to build go package!" && exit 1; fi
sudo mv "$PATH_TO_RES/timed-notify" "$BIN_PATH/timed-notify"
if [ $? != 0 ]; then echo "Failed to build to path!" && exit 1; fi

# Copy Resources
echo "Copying Resources to '$RES'"
sudo cp "$PATH_TO_RES/Notification.png" "$RES/"
if [ $? != 0 ]; then echo "Failed to copy resources!" && exit 1; fi

echo "timed-notify Successfully installed! 🚀"
#/bin/sh

# Default Install Path (LINUX).
RES='/usr/share/timed-notify'
BIN_PATH="/usr/bin"

# MacOS
if [ `uname | grep Darwin` == "Darwin" ]; then
  echo "MacOS Install"
  BIN_PATH="/usr/local/bin"
  RES='/usr/local/share/timed-notify'
fi

# Remove Binary and Resources
echo "Removing Resources Directory '$RES'."
sudo rm -rfv "$RES"
if [ $? != 0 ]; then echo "Failed to remove resources!" && exit 1; fi

echo "Removing Binary from to /usr/bin."
sudo rm -v "$BIN_PATH/timed-notify"
if [ $? != 0 ]; then echo "Failed to remove binary!" && exit 1; fi

echo "timed-notify Successfully uninstalled! 😢"
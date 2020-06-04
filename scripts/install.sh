#/bin/sh

RES='/usr/share/timed-notify'

# Populate Directory
echo "Creating Resources Directory '$RES'"
sudo mkdir -p $RES/res

# Build Package
echo "Building Package..."
cd .. && make build

# Copy Binary and Resources
echo "Copying Resources to '$RES/res/'"
sudo cp ./bin/Notification.png $RES/res/

echo "Copying Binary Over to /usr/bin"
sudo cp ./bin/timed-notify /usr/bin/

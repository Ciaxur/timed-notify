#/bin/sh

RES='/usr/share/timed-notify'


# Remove Binary and Resources
echo "Removing Resources Directory '$RES'"
sudo rm -rfv $RES

echo "Removing Binary from to /usr/bin"
sudo rm -v /usr/bin/timed-notify

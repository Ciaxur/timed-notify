APP = timed-notify

# Build Source
# Copy Over Default Notifcation Image
# TODO: Dynamically set version into the version command by statically
# passing it into the build from branch release name r"^v(\d+\.){2}(\d+)$" else
# "rev hash".
sed -i "s/BuildDate.*/BuildDate=\"$(date "+%Y\/%m\/%d")\"/" internal/build/build.go

build:
	go build -o ./bin/$(APP) ./cmd/timed-notify
	cp ./res/Notification.png ./bin

# Clean up Dev Bin Directory
clean:
	rm -rf ./bin/*


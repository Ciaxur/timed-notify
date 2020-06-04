APP = timed-notify

# Build Source
# Copy Over Default Notifcation Image
build:
	go build -o ./bin/$(APP) ./src
	cp ./res/Notification.png ./bin

# Clean up Dev Bin Directory
clean:
	rm -rf ./bin/*


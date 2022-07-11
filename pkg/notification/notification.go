package notification

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// ReminderStatus -
//  Reminder Time Information.
type ReminderStatus struct {
	RemindIn  time.Duration `json:"remindIn"`
	StartTime time.Time     `json:"startTime"`
	EndTime   time.Time     `json:"endTime"`
}

// NotificationInfo -
//  Information regarding the Notification.
type NotificationInfo struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Icon    string `json:"icon"`
	Urgency string `json:"urgency"`
}

// Notification -
//  Currently Running Process' Status Structure.
type Notification struct {
	PidFilePath  string           `json:"pidFilePath"`
	Pid          int              `json:"pid"`
	Reminder     ReminderStatus   `json:"reminder"`
	Notification NotificationInfo `json:"notification"`
	Version      string           `json:"version"`
}

func NewNotification() *Notification {
	n := &Notification{}
	return n
}

func NewNotificationFromFile(filePath string) (*Notification, error) {
	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file from '%s': %v", filePath, err)
	}

	// Parse JSON Data if no error
	var process *Notification
	process = &Notification{}
	if err := json.Unmarshal(buffer, process); err != nil {
		return nil, fmt.Errorf("failed to parse notification file: %v", err)
	}

	return process, nil
}

func UrgencyIntToString(urgency int) string {
	var urgentStr = "normal"
	switch urgency {
	case 1:
		urgentStr = "low"
	case 2:
		urgentStr = "normal"
	case 3:
		urgentStr = "critical"
	}
	return urgentStr
}

// WriteToFle -
//  Outputs Notification Structure to given File.
// @param file - File to be written to.
func (n *Notification) WriteToFile(file *os.File) {
	buffer, _ := json.MarshalIndent(n, "", "\t")
	file.Write(buffer)
}

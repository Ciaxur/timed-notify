package Management

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

// ReminderStatus -
//  Reminder Time Information
type ReminderStatus struct {
	RemindIn  time.Duration `json:"remindIn"`
	StartTime time.Time     `json:"startTime"`
	EndTime   time.Time     `json:"endTime"`
}

// NotificationInfo -
//  Information regarding the Notification
type NotificationInfo struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Icon    string `json:"icon"`
	Urgency string `json:"urgency"`
}

// ProcessStatus -
//  Currently Running Process' Status Structure
type ProcessStatus struct {
	PidFilePath  string           `json:"pidFilePath"`
	Pid          int              `json:"pid"`
	Reminder     ReminderStatus   `json:"reminder"`
	Notification NotificationInfo `json:"notification"`
	Version      string           `json:"version"`
}

// PidFileWrite -
//  Outputs Process Status Structure to given File
// @param processStatus - ProcessStatus Structure to Write to file
// @param file - File to be written to
func PidFileWrite(processStatus *ProcessStatus, file *os.File) {
	buffer, _ := json.MarshalIndent(processStatus, "", "\t")
	file.Write(buffer)
}

// PidFileRead -
//  Reads Process Status Structure from given File Path
// @param filePath - File to be read from
// @returns processStatus - Structure read from File
// @returns *error - Error for reading File
func PidFileRead(filePath string) (*ProcessStatus, error) {
	buffer, err := ioutil.ReadFile(filePath)
	var processStatus *ProcessStatus = nil

	// Parse JSON Data if no error
	if err == nil {
		processStatus = &ProcessStatus{}
		err = json.Unmarshal(buffer, processStatus)
	}

	return processStatus, err
}

package notification

import (
	"fmt"
	"os"
	"path"

	"github.com/ciaxur/timed-notify/v1/internal/config"
)

// NotificationManager -
//  Stores and handles overall active notification instances.
type NotificationManager struct {
	processes []Notification
}

func NewNotificationManager() (*NotificationManager, error) {
	n := &NotificationManager{
		processes: []Notification{},
	}

	// Parse the active pids.
	if _, err := os.Stat(config.PidDir); !os.IsNotExist(err) {
		pidFiles, _ := os.ReadDir(config.PidDir)

		for _, pid := range pidFiles {
			fullPath := path.Join(config.PidDir, pid.Name())
			process, err := NewNotificationFromFile(fullPath)
			if err != nil {
				return nil, fmt.Errorf("failed to created notification from file: %v", err)
			}
			n.processes = append(n.processes, *process)
		}
	}

	return n, nil
}

func (n *NotificationManager) GetActiveProcesses() []Notification {
	return n.processes
}

func (n *NotificationManager) GetPid(pid int) *Notification {
	for _, process := range n.processes {
		if process.Pid == pid {
			return &process
		}
	}
	return nil
}

//go:build linux

package notification

import (
	"fmt"
	"os/exec"
	"runtime"
)

func (n *Notification) TriggerNotification() error {
	if runtime.GOOS != "linux" {
		return fmt.Errorf("unsupported Platform '%s' 😢", runtime.GOOS)
	}

	cmd := exec.Command("notify-send", n.Notification.Title, n.Notification.Summary, "-i", n.Notification.Icon, "-u", n.Notification.Urgency)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to invoke notify-send command on linux platform: %v", err)
	}
	return nil
}

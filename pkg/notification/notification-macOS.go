//go:build darwin

package notification

import (
	"fmt"

	gosxnotifier "github.com/deckarep/gosx-notifier"
)

func (n *Notification) TriggerNotification() error {
	if runtime.GOOS != "darwin" {
		return fmt.Errorf("unsupported Platform '%s' 😢", runtime.GOOS)
	}

	notify := gosxnotifier.NewNotification(n.Notification.Summary)
	notify.Title = n.Notification.Title
	notify.Sound = gosxnotifier.Tink
	notify.Group = "com.ciaxur.time-notify"
	notify.AppIcon = n.Notification.Icon
	notify.ContentImage = n.Notification.Icon
	if err := notify.Push(); err != nil {
		return fmt.Errorf("failed to push MacOS notification: %v", err)
	}
	return nil
}

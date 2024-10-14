package notification

import (
	"context"
	"os/exec"

	"github.com/deckarep/gosx-notifier"
)

func (n Notification) newSendCommand(ctx context.Context, title, body string) (*exec.Cmd, func(), error) {
	gosxNotification := gosxnotifier.Notification{
		Title:        title,
		Message:      body,
		Link:         n.AppName,
		Sender:       n.AppName,
		AppIcon:      n.Icon,
		ContentImage: n.Icon,
	}

	if n.UrgencyLevel == UrgencyLevelCritical {
		gosxNotification.Sound = gosxnotifier.Glass
	}

	cmd, err := gosxNotification.BuildCommand(ctx)
	if err != nil {
		return nil, nil, err
	}

	return cmd, nil, nil
}

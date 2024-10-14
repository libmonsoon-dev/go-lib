package notification

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"golang.org/x/sys/windows"
	"gopkg.in/toast.v1"
)

// TODO: support for windows < 10
func (n Notification) newSendCommand(ctx context.Context, title, body string) (*exec.Cmd, func(), error) {
	// TODO: support for windows < 10
	maj, _, _ := windows.RtlGetNtVersionNumbers()
	if maj < 10 {
		return nil, nil, fmt.Errorf("not supported")
	}

	toastNotification := &toast.Notification{
		Title:   title,
		Message: body,
		AppID:   n.AppName,
		Icon:    n.Icon,
	}
	if n.UrgencyLevel == UrgencyLevelCritical {
		toastNotification.Audio = toast.LoopingAlarm
	}

	if n.Expire > (3 * time.Second) {
		toastNotification.Duration = toast.Long
	} else {
		toastNotification.Duration = toast.Short
	}

	cmd, err := toastNotification.BuildCommand(ctx)
	cleanup := func() { toastNotification.Cleanup() }
	if err != nil {
		return nil, cleanup, err
	}

	return cmd, cleanup, nil
}

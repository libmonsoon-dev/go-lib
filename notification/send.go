package notification

import (
	"context"
	"fmt"
	"time"

	"github.com/libmonsoon-dev/go-lib/exec"
)

type UrgencyLevel byte

const (
	UrgencyLevelNotDefined UrgencyLevel = iota
	UrgencyLevelLow
	UrgencyLevelNormal
	UrgencyLevelCritical
)

type Notification struct {
	UrgencyLevel UrgencyLevel
	Expire       time.Duration
	AppName      string
	Icon         string
	Category     []string
	Hint         map[string]any
}

func Send(ctx context.Context, title, body string) error {
	var n Notification
	return n.Send(ctx, title, body)
}

func (n Notification) Send(ctx context.Context, title, body string) error {
	cmd, cleanup, err := n.newSendCommand(ctx, title, body)
	if cleanup != nil {
		defer cleanup()
	}
	if err != nil {
		return fmt.Errorf("build command: %w", err)
	}

	_, _, err = exec.RunCommand(cmd)
	return err
}

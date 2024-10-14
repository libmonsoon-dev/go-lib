package notification

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
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

	output, err := cmd.CombinedOutput()
	if err != nil {
		args := append([]string{cmd.Args[0]}, quoteSlice(cmd.Args[1:])...)
		if len(output) > 0 {
			return fmt.Errorf("execute %s: %s: %w", strings.Join(args, " "), string(output), err)
		}
		return fmt.Errorf("execute %s: %w", strings.Join(args, " "), err)
	}
	return err
}

func quoteSlice(args []string) []string {
	result := make([]string, len(args))

	for i := range args {
		result[i] = strconv.Quote(args[i])
	}

	return result
}

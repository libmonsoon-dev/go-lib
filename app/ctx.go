package app

import (
	"context"
	"os/signal"
	"syscall"
)

func Context() (context.Context, context.CancelFunc) {
	ctx, stopNotify := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	return ctx, stopNotify
}

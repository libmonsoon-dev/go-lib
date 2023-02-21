package mainutils

import (
	"context"
	"errors"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/libmonsoon-dev/go-lib/errutils"
)

func init() {
	ctx, cancel = NotifyContext(ctx)
	group, ctx = errgroup.WithContext(ctx)
}

// NotifyContext return context that is marked done
// (its Done channel is closed) when one of the syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP signals arrives,
// when the returned stop function is called, or when the parent context's
// Done channel is closed, whichever happens first.
// If first argument is nil context.Background() will be used
func NotifyContext(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}

	return signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
}

func Context() context.Context {
	return ctx
}

var errs = make(chan error, 1)

func Go(name string, fn func(context.Context) error) {
	runningJobs.Add(name)

	group.Go(func() error {
		err := errutils.Wrap("background job "+name, fn(ctx))
		runningJobs.Remove(name)

		if IgnoreContextError && isContextError(err) {
			return err
		}

		errs <- err
		return err
	})
}

func isContextError(err error) bool {
	return ctx.Err() != nil && errors.Is(err, ctx.Err())
}

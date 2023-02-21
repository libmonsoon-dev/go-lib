package mainutils

import (
	"context"
	"errors"
	"fmt"
	"github.com/libmonsoon-dev/go-lib/async"
	"strings"
	"time"

	"github.com/libmonsoon-dev/go-lib/async/errgroup"
)

var TerminationTimeout = 5 * time.Second

var (
	ctx    context.Context
	cancel context.CancelFunc
	group  *errgroup.Group

	runningJobs = async.NewSet[string]()
)

func waitBackgroundJobs(err *error) {
	cancel()

	oneOf := make(chan struct{}, 2)
	go func() {
		time.Sleep(TerminationTimeout)

		addError(err, context.Cause(ctx))
		running := strings.Join(runningJobs.Values(), ", ")
		addError(err, fmt.Errorf("termination %w: still running: [%v]", context.DeadlineExceeded, running))

		oneOf <- struct{}{}
	}()

	go func() {
		addError(err, group.Wait())

		oneOf <- struct{}{}
	}()

	<-oneOf
	return
}

func addError(out *error, err error) {
	if errors.Is(err, context.Canceled) {
		return
	}

	*out = errors.Join(*out, err)
}

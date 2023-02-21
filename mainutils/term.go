package mainutils

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/libmonsoon-dev/go-lib/async"
)

var (
	TerminationTimeout = 5 * time.Second
	IgnoreContextError = true
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	group  *errgroup.Group

	runningJobs = async.NewSet[string]()
)

func terminateBackgroundJobs(err *error) {
	cancel()

	oneOf := make(chan struct{}, 2)
	go func() {
		time.Sleep(TerminationTimeout)

		addBackgroundErrors(err)
		running := strings.Join(runningJobs.Values(), ", ")
		addError(err, fmt.Errorf("termination %w: still running: [%v]", context.DeadlineExceeded, running))

		oneOf <- struct{}{}
	}()

	go func() {
		group.Wait()
		addBackgroundErrors(err)
		oneOf <- struct{}{}
	}()

	<-oneOf
	return
}

func addBackgroundErrors(out *error) {
	for {
		select {
		case err := <-errs:
			addError(out, err)
		default:
			return
		}
	}
}

func addError(out *error, err error) {
	*out = errors.Join(*out, err)
}

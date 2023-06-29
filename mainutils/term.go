package mainutils

import (
	"context"
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
		_ = group.Wait()
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
	*out = errorsJoin(*out, err)
}

// TODO: remove afer upgrade to go 1.20
func errorsJoin(errs ...error) error {
	n := 0
	for _, err := range errs {
		if err != nil {
			n++
		}
	}

	if n == 0 {
		return nil
	}

	e := &joinError{
		errs: make([]error, 0, n),
	}

	for _, err := range errs {
		if err != nil {
			e.errs = append(e.errs, err)
		}
	}

	return e
}

type joinError struct {
	errs []error
}

func (e *joinError) Error() string {
	var b []byte
	for i, err := range e.errs {
		if i > 0 {
			b = append(b, '\n')
		}

		b = append(b, err.Error()...)
	}
	return string(b)
}

func (e *joinError) Unwrap() []error {
	return e.errs
}

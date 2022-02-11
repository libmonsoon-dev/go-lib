package call

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/libmonsoon-dev/go-lib/async"
	"github.com/libmonsoon-dev/go-lib/builtintools"
	"github.com/libmonsoon-dev/go-lib/errutils"
)

type GroupConfig struct {
	Context context.Context
	MaxJobs int
}

func NewGroup(conf GroupConfig) *Group {
	g := &Group{}
	return g.
		SetMaxJobs(conf.MaxJobs).
		SetContext(conf.Context)
}

type Group struct {
	resultLock      sync.Mutex
	args            Args
	lastResultIndex async.IdGenerator

	ctx   context.Context
	group errgroup.Group

	semaphore      async.Semaphore
	doOrMakeCalled bool
}

func (g *Group) SetMaxJobs(maxJobs int) *Group {
	if maxJobs > 0 {
		g.semaphore = async.NewSemaphore(maxJobs)
	} else {
		g.semaphore = nil
	}

	return g
}

// SetContext Sets ctx and makes the inner init based on it.
// The method panics if .Do or .Make was called
func (g *Group) SetContext(ctx context.Context) *Group {
	if g.doOrMakeCalled {
		panic("call.Group.SetContext: illigal usage")
	}
	if ctx == nil {
		g.group = errgroup.Group{}
	} else {
		var group *errgroup.Group
		group, g.ctx = errgroup.WithContext(ctx)

		// legal copy before first use
		//nolint:govet
		g.group = *group

	}

	return g
}

func (g *Group) GetContext() context.Context {
	if g.ctx == nil {
		return context.Background()
	}
	return g.ctx
}

func (g *Group) Do(errorMessage string, fn func() error) *Group {
	g.group.Go(func() error {
		g.acquireSemaphore()
		defer g.releaseSemaphore()

		return errutils.Wrap(errorMessage, fn())
	})

	return g
}

func (g *Group) Make(errorMessage string, fn func() (any, error)) *Group {
	index := g.lastResultIndex.Generate()

	g.Do(errorMessage, func() error {
		result, err := fn()

		g.resultLock.Lock()
		defer g.resultLock.Unlock()

		if g.args == nil {
			g.args = *builtintools.AcquireAnySlice()
		}
		g.args.Grow(int(index) + 1)
		g.args[index] = result

		return err
	})

	return g
}

func (g *Group) DoChain(errorMessage string, fn func(*Chain)) *Group {
	g.doOrMakeCalled = true
	g.Do(errorMessage, func() error {
		chain := AcquireChain()
		defer ReleaseChain(chain)

		fn(chain)
		return chain.GetError()
	})

	return g
}

func (g *Group) MakeChain(errorMessage string, fn func(*Chain)) *Group {
	g.doOrMakeCalled = true
	g.Make(errorMessage, func() (any, error) {
		chain := AcquireChain()
		defer ReleaseChain(chain)

		fn(chain)
		return chain.GetResult()
	})

	return g
}

func (g *Group) GetArgs() Args {
	g.resultLock.Lock()
	defer g.resultLock.Unlock()

	return g.args
}

func (g *Group) SetArgs(args Args) {
	g.resultLock.Lock()
	defer g.resultLock.Unlock()

	g.args = args
}

func (g *Group) GetError() error {
	return g.group.Wait()
}

func (g *Group) Result() (args Args, err error) {
	err = g.GetError()
	args = g.GetArgs()

	return
}

func (g *Group) Reset() {
	g.args.Reset()
	g.lastResultIndex.Reset()

	g.ctx = nil
	g.group = errgroup.Group{}

	g.semaphore = nil
	g.doOrMakeCalled = false
}

func (g *Group) acquireSemaphore() {
	if g.semaphore == nil {
		return
	}

	g.semaphore.Acquire()
}

func (g *Group) releaseSemaphore() {
	if g.semaphore == nil {
		return
	}

	g.semaphore.Release()
}

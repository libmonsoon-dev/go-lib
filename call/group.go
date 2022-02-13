package call

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/libmonsoon-dev/go-lib/async"
	"github.com/libmonsoon-dev/go-lib/builtintools"
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

func (g *Group) DoFunc(errorMessage string, fn DoFunc) Manager {
	return g.Do(DoArgs{ErrorMessage: errorMessage, Func: fn})
}

func (g *Group) MakeFunc(errorMessage string, fn MakeFunc[any]) Manager {
	return g.Make(MakeArgs[any]{ErrorMessage: errorMessage, Func: fn})
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
// The method panics if .Do* or .Make* was called
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

func (g *Group) Do(args DoArgs) Manager {
	g.group.Go(func() error {
		g.acquireSemaphore()
		defer g.releaseSemaphore()

		err := g.getDoFunc(args)()
		if err != nil {
			return fmt.Errorf(args.GetErrorMessage()+": %w", err)
		}

		return nil
	})

	return g
}

func (g *Group) Make(args MakeArgs[any]) Manager {
	index := g.lastResultIndex.Generate()

	doArgs := DoArgs{
		Func: func() error {
			result, err := g.getMakeFunc(args)()

			g.resultLock.Lock()
			defer g.resultLock.Unlock()

			if g.args == nil {
				g.args = *builtintools.AcquireAnySlice()
			}
			g.args.Grow(int(index) + 1)
			g.args[index] = result

			return err
		},

		ErrorMessage:     args.ErrorMessage,
		ErrorMessageFunc: args.ErrorMessageFunc,
	}

	g.Do(doArgs)

	return g
}

func (g *Group) GetArgs() Args {
	g.resultLock.Lock()
	defer g.resultLock.Unlock()

	return g.args
}

func (g *Group) SetArgs(args Args) Manager {
	g.resultLock.Lock()
	defer g.resultLock.Unlock()

	g.args = args
	return g
}

func (g *Group) GetError() error {
	return g.group.Wait()
}

func (g *Group) GetResult() (args Args, err error) {
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

func (g *Group) getDoFunc(args DoArgs) (fn DoFunc) {
	fn = args.GetFunc()
	if fn != nil {
		return
	}

	if args.GroupFunc != nil {
		return func() error {
			args.GroupFunc(g)
			return g.GetError()
		}
	}

	if args.ChainFunc != nil {
		return func() (err error) {
			chain := AcquireChain()
			defer ReleaseChain(chain)

			args.ChainFunc(chain)
			return chain.GetError()
		}
	}

	panic("Invalid args")
}

func (g *Group) getMakeFunc(args MakeArgs[any]) (fn MakeFunc[any]) {
	fn = args.GetFunc()
	if fn != nil {
		return
	}

	if args.GroupFunc != nil {
		return func() (any, error) {
			args.GroupFunc(g)
			return g.GetResult()
		}
	}

	if args.ChainFunc != nil {
		return func() (any, error) {
			chain := AcquireChain()
			defer ReleaseChain(chain)

			args.ChainFunc(chain)
			result, err := chain.GetResult()
			return result.Copy(), err
		}

	}

	panic("Invalid args")
}

var _ Manager = (*Group)(nil)

type GroupFunc func(*Group)

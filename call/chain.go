package call

import (
	"fmt"

	"github.com/libmonsoon-dev/go-lib/builtintools"
)

type ChainConfig struct {
	Args Args
}

func NewChain(conf ChainConfig) *Chain {
	conf.Args = *builtintools.AcquireAnySlice()

	return &Chain{
		args: conf.Args,
	}
}

type Chain struct {
	args Args
	err  error
}

func (c *Chain) Do(args DoArgs) Manager {
	if c.err != nil {
		return c
	}

	err := c.getDoFunc(args)()
	if err != nil {
		c.err = fmt.Errorf(args.GetErrorMessage()+": %w", err)
	}

	return c
}

func (c *Chain) Make(args MakeArgs[any]) Manager {
	if c.err != nil {
		return c
	}

	result, err := c.getMakeFunc(args)()
	if err != nil {
		c.err = fmt.Errorf(args.GetErrorMessage()+": %w", err)
	}
	c.args = append(c.args, result)

	return c
}

func (c *Chain) DoFunc(errorMessage string, fn DoFunc) Manager {
	return c.Do(DoArgs{ErrorMessage: errorMessage, Func: fn})
}

func (c *Chain) MakeFunc(errorMessage string, fn MakeFunc[any]) Manager {
	return c.Make(MakeArgs[any]{ErrorMessage: errorMessage, Func: fn})
}

func (c *Chain) GetArgs() Args {
	return c.args
}

func (c *Chain) SetArgs(args Args) Manager {
	c.args = args

	return c
}

func (c *Chain) GetError() error {
	return c.err
}

func (c *Chain) GetResult() (Args, error) {
	return c.GetArgs(), c.GetError()
}

func (c *Chain) Reset() {
	c.args.Reset()
	c.err = nil
}

func (c *Chain) getDoFunc(args DoArgs) (fn DoFunc) {
	fn = args.GetFunc()
	if fn != nil {
		return
	}

	if args.GroupFunc != nil {
		return func() (err error) {
			group := AcquireGroup()
			defer ReleaseGroup(group)

			args.GroupFunc(group)
			return group.GetError()
		}
	}

	if args.ChainFunc != nil {
		return func() error {
			args.ChainFunc(c)
			return c.GetError()
		}
	}

	panic("Invalid args")
}

func (c *Chain) getMakeFunc(args MakeArgs[any]) (fn MakeFunc[any]) {
	fn = args.GetFunc()
	if fn != nil {
		return
	}

	if args.GroupFunc != nil {
		return func() (any, error) {
			group := AcquireGroup()
			defer ReleaseGroup(group)

			args.GroupFunc(group)
			result, err := group.GetResult()
			return result.Copy(), err
		}
	}

	if args.ChainFunc != nil {
		return func() (any, error) {
			args.ChainFunc(c)
			return c.GetResult()
		}
	}

	panic("Invalid args")
}

var _ Manager = (*Chain)(nil)

type ChainFunc func(*Chain)

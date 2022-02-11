package call

import (
	"fmt"
	"github.com/libmonsoon-dev/go-lib/builtintools"
)

type ChainConfig struct {
	Args Args
}

func NewChain(conf ChainConfig) *Chain {
	conf.Args = builtintools.AcquireAnySlice()

	return &Chain{
		args: conf.Args,
	}
}

type Chain struct {
	args Args
	err  error
}

func (c *Chain) Do(errorMessage string, fn func() error) *Chain {
	if c.err != nil {
		return c
	}

	if c.err = fn(); c.err != nil {
		c.err = fmt.Errorf(errorMessage+": %w", c.err) // TODO: errors wrap method that handle input nil error
	}

	return c
}

func (c *Chain) Make(errorMessage string, fn func() (any, error)) *Chain {
	if c.err != nil {
		return c
	}

	result, err := fn()
	if err != nil {
		c.err = fmt.Errorf(errorMessage+": %w", err)
	}
	c.args = append(c.args, result)

	return c
}

func (c *Chain) DoGroup(errorMessage string, fn func(*Group)) *Chain {
	c.Do(errorMessage, func() (err error) {
		group := AcquireGroup()
		defer ReleaseGroup(group)

		fn(group)
		return group.GetError()
	})

	return c
}

func (c *Chain) MakeGroup(errorMessage string, fn func(*Group)) *Chain {
	c.Make(errorMessage, func() (any, error) {
		group := AcquireGroup()
		defer ReleaseGroup(group)

		fn(group)
		result, err := group.Result()
		return result.Copy(), err
	})

	return c
}

func (c *Chain) GetArgs() Args {
	return c.args
}

func (c *Chain) SetArgs(args Args) *Chain {
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

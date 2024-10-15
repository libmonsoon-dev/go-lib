package context

import (
	"context"
	"time"

	"github.com/libmonsoon-dev/go-lib/cgo-friendly/channel"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() channel.ReadOnly[struct{}]
	Err() error
	Value(key any) any
}

var _ Context = wrapper{}

type wrapper struct {
	ctx context.Context
}

func (w wrapper) Done() channel.ReadOnly[struct{}] {
	return channel.NewReadOnly[struct{}](w.ctx.Done())
}

func (w wrapper) Deadline() (deadline time.Time, ok bool) {
	return w.ctx.Deadline()
}

func (w wrapper) Err() error {
	return w.ctx.Err()
}

func (w wrapper) Value(key any) any {
	return w.ctx.Value(key)
}

func Wrap(ctx context.Context) Context {
	return wrapper{ctx: ctx}
}

func Background() Context {
	return Wrap(context.Background())
}

func TODO() Context {
	return Wrap(context.TODO())
}

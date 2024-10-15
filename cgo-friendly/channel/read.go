package channel

func NewReadOnly[T any](ch <-chan T) ReadOnly[T] {
	return ReadOnly[T]{
		ch: ch,
	}
}

type ReadOnly[T any] struct {
	ch <-chan T
}

// Recv receives and returns a value from the channel
// The receive blocks until a value is ready.
// The boolean value ok is true if the value x corresponds to a send
// on the channel, false if it is a zero value received because the channel is closed.
func (c ReadOnly[T]) Recv() (v T, ok bool) {
	v, ok = <-c.ch
	return v, ok
}

// TryRecv attempts to receive a value from the channel but will not block
// If the receive delivers a value, v is the transferred value and ok is true
// If the receive cannot finish without blocking, v is the zero value and ok is false
// If the channel is closed, v is the zero value for the channel's element type and ok is false.
func (c ReadOnly[T]) TryRecv() (v T, ok bool) {
	select {
	case v, ok = <-c.ch:
	// pass
	default:
		// pass
	}

	return v, ok
}

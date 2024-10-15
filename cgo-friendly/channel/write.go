package channel

func NewWriteOnly[T any](ch chan<- T) WriteOnly[T] {
	return WriteOnly[T]{
		ch: ch,
	}
}

type WriteOnly[T any] struct {
	ch chan<- T
}

// Send sends v on the channel
func (c WriteOnly[T]) Send(v T) {
	c.ch <- v
}

// TrySend attempts to send v on the channel but will not block
// It reports whether the value was sent
func (c WriteOnly[T]) TrySend(v T) (ok bool) {
	select {
	case c.ch <- v:
		ok = true
	default:
		// ok = false
	}

	return ok
}

// Close closes the channel
func (c WriteOnly[T]) Close() {
	close(c.ch)
}

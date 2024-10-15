package channel

func New[T any]() Channel[T] {
	return Make[T](0)
}

func Make[T any](size int) Channel[T] {
	ch := make(chan T, size)

	return Channel[T]{
		ReadOnly:  ReadOnly[T]{ch: ch},
		WriteOnly: WriteOnly[T]{ch: ch},
	}
}

type Channel[T any] struct {
	ReadOnly[T]
	WriteOnly[T]
}

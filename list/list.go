package list

func NewList[T any](args ...T) *List[T] {
	l := &List[T]{}

	for _, val := range args {
		l.Append(val)
	}

	return l
}

type List[T any] struct {
	First *Node[T]
	Last  *Node[T]
}

type Node[T any] struct {
	Next  *Node[T]
	Value T
}

func (l *List[T]) Append(value T) {
	next := &Node[T]{Value: value}
	if l.Last != nil {
		l.Last.Next = next
	}
	l.Last = next

	if l.First == nil {
		l.First = next
	}
}

func (l *List[T]) IsEmpty() bool {
	return l.Last == nil
}

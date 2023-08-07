package list

func NewList[T any]() *List[T] {
	return &List[T]{}
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

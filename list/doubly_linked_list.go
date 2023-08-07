package list

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

type DoublyLinkedList[T any] struct {
	First *DoublyLinkedNode[T]
	Last  *DoublyLinkedNode[T]
}

type DoublyLinkedNode[T any] struct {
	Next  *DoublyLinkedNode[T]
	Prev  *DoublyLinkedNode[T]
	Value T
}

func (l *DoublyLinkedList[T]) Append(value T) {
	next := &DoublyLinkedNode[T]{Prev: l.Last, Value: value}
	if l.Last != nil {
		l.Last.Next = next
	}
	l.Last = next

	if l.First == nil {
		l.First = next
	}
}

func (l *DoublyLinkedList[T]) Pop() (value T, ok bool) {
	if l.Last == nil {
		return
	}

	last := l.Last
	if last.Prev != nil {
		l.Last = last.Prev
		last.Prev.Next = nil
	} else {
		l.First = nil
		l.Last = nil
	}

	return last.Value, true
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.Last == nil
}

package tree

func NewBinaryTree[T any](values ...T) *BinaryNode[T] {
	root := new(BinaryNode[T])

	if len(values) == 0 {
		return root
	}

	return insertLevelOrder(values, 0)
}

type BinaryNode[T any] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func insertLevelOrder[T any](values []T, i int) (root *BinaryNode[T]) {
	if i < len(values) {
		root = &BinaryNode[T]{
			Value: values[i],
			Left:  insertLevelOrder(values, 2*i+1),
			Right: insertLevelOrder(values, 2*i+2),
		}
	}

	return
}

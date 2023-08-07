package tree

import (
	"reflect"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	want := &BinaryNode[int]{
		Value: 1,
		Left: &BinaryNode[int]{
			Value: 2,
			Left: &BinaryNode[int]{
				Value: 4,
			},
			Right: &BinaryNode[int]{
				Value: 5,
			},
		},
		Right: &BinaryNode[int]{
			Value: 3,
			Left: &BinaryNode[int]{
				Value: 6,
			},
		},
	}
	if got := NewBinaryTree(1, 2, 3, 4, 5, 6); !reflect.DeepEqual(got, want) {
		t.Errorf("NewBinaryTree() = %v, want %v", got, want)
	}
}

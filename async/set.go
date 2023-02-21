package async

import (
	"fmt"
	"sync"

	"golang.org/x/exp/maps"
)

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

type Set[T comparable] struct {
	m  map[T]struct{}
	mu sync.Mutex
}

func (s *Set[T]) Add(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[val] = struct{}{}
}

func (s *Set[T]) Remove(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.m, val)
}

func (s *Set[T]) Values() []T {
	s.mu.Lock()
	defer s.mu.Unlock()

	return maps.Keys(s.m)
}

func (s *Set[T]) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	return fmt.Sprint(s.m)
}

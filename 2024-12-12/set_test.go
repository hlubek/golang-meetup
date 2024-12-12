package main

import (
	"iter"
	"testing"
)

func NewSet[E comparable]() *Set[E] {
	return &Set[E]{m: make(map[E]struct{})}
}

func (s *Set[E]) Add(v E) {
	s.m[v] = struct{}{}
}

// START OMIT

type Set[E comparable] struct{ m map[E]struct{} }

func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}

func BenchmarkRangeOverFunc(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
	var sum int
	for item := range s.All() {
		sum += item
	}
	_ = sum
}

func BenchmarkWithoutIter(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
	var sum int
	for item := range s.m {
		sum += item
	}
	_ = sum
}

// END OMIT

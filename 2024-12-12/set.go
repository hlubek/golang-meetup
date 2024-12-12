package main

import "iter"

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

func main() {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(1)
	for value := range s.All() {
		println(value)
	}
}

// END OMIT

package main

import (
	"cmp"
	"iter"
	"slices"
)

// START OMIT

func MergeSorted[E cmp.Ordered](s1, s2 iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		next1, stop1 := iter.Pull(s1)
		defer stop1()
		next2, stop2 := iter.Pull(s2)
		defer stop2()
		v1, ok1 := next1()
		v2, ok2 := next2()

		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 < v2) {
				if !yield(v1) {
					return
				}
				v1, ok1 = next1()
			} else {
				if !yield(v2) {
					return
				}
				v2, ok2 = next2()
			}
		}
	}
}

// END OMIT

func main() {
	s1 := slices.Values([]string{"apple", "ashberry", "banana", "cherry"})
	s2 := slices.Values([]string{"cucumber", "pear", "watermelon", "zitrone"})

	for value := range MergeSorted(s1, s2) {
		println(value)
	}
}

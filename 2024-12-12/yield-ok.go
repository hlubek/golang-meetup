package main

import "iter"

func MyIterator() iter.Seq[int] {
	return func(yield func(int) bool) {
		if !yield(1) {
			return
		}
		if !yield(2) {
			return
		}
	}
}

func main() {
	for value := range MyIterator() {
		println(value)
		break // Return after first value
	}
}

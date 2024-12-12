package main

import "iter"

func MyIterator() iter.Seq[int] {
	return func(yield func(int) bool) {
		yield(1)
	}
}

func main() {
	for value := range MyIterator() {
		println(value)
	}
}

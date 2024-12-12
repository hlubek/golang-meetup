package main

import "iter"

func MyIterator() iter.Seq[int] {
	return func(yield func(int) bool) {
		yield(1)
		yield(2)
	}
}

func main() {
	for value := range MyIterator() {
		println(value)
		break // Let's break after the first value
	}
}

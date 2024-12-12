package main

import "iter"

// START OMIT

func Fibonacci() iter.Seq2[int, int] {
	i, j := 0, 1
	n := 0
	return func(yield func(int, int) bool) {
		for {
			if !yield(n, i) {
				return
			}
			i, j = j, i+j
			n++
		}
	}
}

func main() {
	for n, value := range Fibonacci() {
		println(n, value)
		if n >= 10 {
			break
		}
	}
}

// END OMIT

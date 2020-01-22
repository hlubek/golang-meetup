package main

import "fmt"

// START OMIT

func process(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4}
	var factor int
	double := func(i int) int {
		return i * factor
	}
	factor = 2
	out := process(numbers, double)

	fmt.Println(out)
}

// END OMIT

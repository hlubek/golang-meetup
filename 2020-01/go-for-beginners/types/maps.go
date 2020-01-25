package main

import "fmt"

func main() {
	// START OMIT
	fruits := []string{"apple", "banana", "kiwi", "orange", "banana", "kiwi", "banana"}

	m := make(map[string]int)
	for _, fruit := range fruits {
		m[fruit]++
	}

	for fruit := range m {
		fmt.Printf("%ss: %d\n", fruit, m[fruit])
	}
	// END OMIT
}

package main

import "fmt"

func main() {
	// START OMIT
	// Map literal
	m := map[string]int{
		"apple":  0,
		"banana": 4,
		"kiwi":   1,
	}

	// Gets empty value
	pears := m["pear"]
	fmt.Printf("%d pears", pears)

	// Special form to check if element exists
	cherries, ok := m["cherry"]
	if ok {
		fmt.Printf("%d cherries", cherries)
	}

	// END OMIT
}

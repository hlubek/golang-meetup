package main

import "fmt"

func main() {
	// START ARRAY OMIT
	loc := [2]float64{54.32133, 10.13489}
	// END ARRAY OMIT

	// START SLICE OMIT
	names := []string{"John", "Jane"}
	names = append(names, "Dave")
	// END SLICE OMIT

	// START SLICE EXPR OMIT
	lastTwo := names[1:3]

	fmt.Println(lastTwo)
	// END SLICE EXPR OMIT

	_ = loc
	_ = names
	_ = lastTwo
}

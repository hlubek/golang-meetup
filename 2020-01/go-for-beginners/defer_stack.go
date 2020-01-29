package main

import "fmt"

// START OMIT
func a() {
	defer fmt.Println("A defer")
	fmt.Println("In A")
}

func main() {
	fmt.Println("In main")

	defer fmt.Println("Main defer #1")
	defer fmt.Println("Main defer #2")

	a()
}

// END OMIT

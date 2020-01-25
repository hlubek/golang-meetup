package main

import "fmt"

func main() {
	// START OMIT
	i := 1

	// j is a pointer to i
	j := &i

	// k is a _copy_ of i
	k := i

	i++

	// *j dereferences a pointer
	fmt.Println(i, *j, k)
	// END OMIT
}

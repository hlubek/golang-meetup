package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	go func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Wake up!")
	}()

	go func() {
		time.Sleep(5 * time.Millisecond)
		fmt.Println("Get up!")
	}()

	fmt.Println("Wait a second...")

	time.Sleep(1 * time.Second)
	// END OMIT
}

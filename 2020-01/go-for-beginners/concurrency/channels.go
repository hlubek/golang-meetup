package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	c := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		c <- "Wake up!"
	}()

	go func() {
		time.Sleep(5 * time.Millisecond)
		c <- "Get up!"
	}()

	fmt.Println("Wait for it...")
	fmt.Println(<-c)
	fmt.Println(<-c)

	// END OMIT
}

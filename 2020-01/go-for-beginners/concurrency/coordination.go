package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// START OMIT
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Wake up!")

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		time.Sleep(5 * time.Millisecond)
		fmt.Println("Get up!")

		wg.Done()
	}()

	fmt.Println("Wait for it...")
	wg.Wait()
	// END OMIT
}

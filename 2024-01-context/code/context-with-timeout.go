package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	neverReady := make(chan bool)
	shortDuration := 1 * time.Second
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	// START OMIT
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration) // HL
	// Although the context is canceled call the cancel function to release resources
	// associated with the context 
	defer cancel()

	go someSlowRoutine(ctx, neverReady)

	select {
	case <-neverReady:
		fmt.Println("ready")
	case <-ctx.Done(): // HL
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
	// END OMIT
}

func someSlowRoutine(ctx context.Context, ready chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Print("finished running")
	ready <- true
}

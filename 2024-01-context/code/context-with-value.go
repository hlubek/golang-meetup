package main

import (
	"context"
	"errors"
	"fmt"
)

type contextKey int

const requestIDKey contextKey = iota

func main() {
	// START OMIT
	// this is created for each request received by the server
	parentCtx := context.Background()

	// this is executed by middleware A
	ctx := context.WithValue(parentCtx, requestIDKey, "some random identifier") // HL

	middlewareB(ctx)

}

func middlewareB(ctx context.Context) {
	var requestID string
	if v := ctx.Value(requestIDKey); v != nil { // HL
		requestID = v.(string)
	}
	// error handling omitted

	fmt.Println(requestID)
}

// END OMIT

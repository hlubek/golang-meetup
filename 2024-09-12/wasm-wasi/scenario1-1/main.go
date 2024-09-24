package main

import (
	"fmt"
	"net/http"

	"github.com/stealthrocket/net/wasip1"
)

func main() {
	listener, err := wasip1.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	count := 0

	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			count++

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Hello, World! Request count: %d", count)
		}),
	}
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}

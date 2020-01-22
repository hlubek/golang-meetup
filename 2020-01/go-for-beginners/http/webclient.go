package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://localhost:4321")
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}

	io.Copy(os.Stdout, res.Body)
}

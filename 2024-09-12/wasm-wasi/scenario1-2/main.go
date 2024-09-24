package main

import (
	"io"
	"net/http"
	"os"

	_ "github.com/stealthrocket/net/http"
)

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

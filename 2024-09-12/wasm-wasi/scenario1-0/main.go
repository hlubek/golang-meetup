package main

import "os"

func main() {
	d, err := os.ReadDir("/")
	if err != nil {
		panic(err)
	}
	for _, de := range d {
		println("-", de.Name())
	}
	src, err := os.ReadFile("/main.wasm")
	if err != nil {
		panic(err)
	}
	println("main.wasm:", len(src), "bytes")
}

package main

import "fmt"

func main() {
	fmt.Println(fib(5))
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

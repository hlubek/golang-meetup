package main

// START OMIT
import "testing"

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func TestFib(t *testing.T) {
	n := 1
	actual := fib(n)
	expected := 1
	if actual != expected {
		t.Errorf("fib(%d), expected: %d, got: %d", n, expected, actual)
	}
}

// END OMIT

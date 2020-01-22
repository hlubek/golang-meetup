package main

import "errors"

func do(input string) (int, error) {
	if input == "" {
		return 0, errors.New("empty input")
	}
	return len(input), nil
}

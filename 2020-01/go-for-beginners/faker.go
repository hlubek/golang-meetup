package main

import (
	"fmt"

	"github.com/bxcodec/faker"
)

func main() {
	fmt.Printf("Hello %s", faker.Name())
}

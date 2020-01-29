package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// START OMIT
func main() {
	data, err := readFile("README.md")
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}

	fmt.Println(string(data[:100]))
}

func readFile(name string) ([]byte, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close() // HLdefer

	return ioutil.ReadAll(f)
}

// END OMIT

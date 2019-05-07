package main

import (
	"fmt"
	"os"
)

func main() {
	// initialized with empty string
	var s, sep string
	// how to get argument from command line, start from index 1
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] // concatenation
		sep = " "
	}
	fmt.Println(s)
}

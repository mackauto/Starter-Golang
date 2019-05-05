package main

import (
	"fmt"
	"os"
)

// how to get argument from command line
func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ":"
	}
	fmt.Println(s)
}

func main() {
	echo1()
}

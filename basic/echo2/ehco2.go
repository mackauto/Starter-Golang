package main

import (
	"fmt"
	"os"
)

func main() {
	// short variable declaration, better use within function
	s, sep := "", ""
	// how to use range and slice
	// index start from 0 and omit with _
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

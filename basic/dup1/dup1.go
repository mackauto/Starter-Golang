package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// how to create dictionary/map
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// how to read from os standard input
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line) // how to format output
			/*
			%d for decimal integer
			%x for integer in hexadecimal
			%o for integer in octal
			%b for integer in binary
			%f %g %e for floating point number
			%t for boolean
			%c for rune/Unicode code point
			%s for string
			%q for quoted string like "abc" or rune 'c'
			%v for any value in a natural format
			%T for type of any value
			%% for % itself
			*/
		}
	}
}

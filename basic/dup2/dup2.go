package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// counts is a map, a map is a reference to the data structure
		// countLines receives a copy of the reference
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// how to open file
			f, err := os.Open(arg)
			if err != nil {
				// how to print err
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			_ = f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// how to define a function
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

package main

import (
	"flag"
	"fmt"
	"strings"
)

// how to create command line option
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	// update flag variable from the default value
	flag.Parse()
	// use pointer to get value of command line option
	fmt.Printf(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

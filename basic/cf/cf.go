package main

// how to use my own package

import (
	"fmt"
	"github.com/JPMike/Starter-Golang/basic/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		// note that %s format use method String() of type Fahrenheit and Celsius
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

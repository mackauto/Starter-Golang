package main

import "fmt"

// package level declaration
const boilingF = 212.0

func main() {
	// local variable to function main
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}

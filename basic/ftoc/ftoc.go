package main

import "fmt"

func main() {
	// how to define constant
	const freezingF, boilingF = 32.0, 212.0
	// how to print with format
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))
}

// how to define function
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

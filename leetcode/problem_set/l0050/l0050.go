package main

import "fmt"

func myPow(x float64, n int) float64 {
	// base condition
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n&1 == 1 {
		return x * myPow(x*x, n/2)
	}
	return myPow(x*x, n/2)
}

func myPowIterationMethod(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	pow := 1.0
	for n > 0 {
		if n&1 == 1 {
			pow *= x
		}
		x *= x
		n = n >> 1
	}
	return pow
}

func main() {
	x1 := 2.0
	n1 := -2
	fmt.Println(myPow(x1, n1))
	fmt.Println(myPowIterationMethod(x1, n1))
}

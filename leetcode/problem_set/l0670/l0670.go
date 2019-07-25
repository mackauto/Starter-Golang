package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	numStr := strconv.Itoa(num)
	numRune := []rune(numStr)
	max := num
	for i := 0; i < len(numRune); i++ {
		for j := i + 1; j < len(numRune); j++ {
			numRune[i], numRune[j] = numRune[j], numRune[i]
			temp, _ := strconv.Atoi(string(numRune))
			if temp > max {
				max = temp
			}
			numRune[i], numRune[j] = numRune[j], numRune[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maximumSwap(2736))
	fmt.Println(maximumSwap(9973))
}

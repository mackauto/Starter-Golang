package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var ans int
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		ans = 0
	}
	return ans
}

func main() {
	nums := []int{123, -123, 120}
	for _, v := range nums {
		fmt.Println(reverse(v))
	}
}

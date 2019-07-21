package main

import (
	"fmt"
	"math"
	"strings"
)

const numToStr = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func to10(num string) int {
	num10 := 0.0
	for i := len(num) - 1; i >= 0; i-- {
		s := string(num[i])
		num10 += float64(strings.Index(numToStr, s)) * math.Pow(float64(36), float64(len(num)-1-i))
	}
	return int(num10)
}

func to36(num int) string {
	s := ""
	for num != 0 {
		i := num % 36
		s = string(numToStr[i]) + s
		num /= 36
	}
	return s
}

func main() {
	fmt.Println(to36(35))
	fmt.Println(to36(36))
	fmt.Println(to36(65036))

	fmt.Println(to10("Z"))
	fmt.Println(to10("10"))
	fmt.Println(to10("1E6K"))
}

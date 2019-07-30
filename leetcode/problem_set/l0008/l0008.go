package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	res, sign, strLen, idx := 0, 1, len(str), 0

	for idx < strLen && (str[idx] == ' ') {
		idx++
	}

	if idx < strLen {
		if str[idx] == '+' {
			sign = 1
			idx++
		} else if str[idx] == '-' {
			sign = -1
			idx++
		}
	}

	for idx < strLen && str[idx] >= '0' && str[idx] <= '9' {
		res = res*10 + int(str[idx]) - '0'
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		} else if sign*res < math.MinInt32 {
			return math.MinInt32
		}
		idx++
	}
	return res * sign
}

func main() {
	strs := []string{"42", "    -42", "4193 with words", "words and 987", "-91283472332"}
	for _, str := range strs {
		fmt.Println(myAtoi(str))
	}
}

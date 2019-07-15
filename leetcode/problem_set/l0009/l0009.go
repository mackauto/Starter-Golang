package main

import "fmt"

func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	return ans
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x != reverse(x) {
		return false
	}
	return true
}

func main() {
	nums := []int{121, -121, 10}
	for _, v := range nums {
		fmt.Println(isPalindrome(v))
	}
}

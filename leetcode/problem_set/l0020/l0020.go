package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if (v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') ||
			(v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	s1 := "()[]{}"
	fmt.Println(isValid(s1))

	s2 := "()[]{}"
	fmt.Println(isValid(s2))

	s3 := "([)]"
	fmt.Println(isValid(s3))
}

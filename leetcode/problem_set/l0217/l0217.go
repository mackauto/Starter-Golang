package main

import "fmt"

func containsDuplicate(nums []int) bool {
	memory := map[int]bool{}
	for _, v := range nums {
		if memory[v] {
			return true
		}
		memory[v] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}

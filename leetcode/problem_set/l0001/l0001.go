package main

import "fmt"

func twoSum(nums []int, target int) []int {
	memory := make(map[int]int)
	var ans []int
	for i, v := range nums {
		if j, ok := memory[target-v]; ok {
			ans = append(ans, j, i)
		} else {
			memory[v] = i
		}
	}
	return ans
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print(twoSum(nums, target))
}

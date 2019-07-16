package main

import (
	"fmt"
	"github.com/JPMike/Starter-Golang/leetcode/util"
)

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func removeDuplicates(nums []int) int {
	memory := make(map[int]bool)
	j := 0
	for i := 0; i < len(nums); i++ {
		if !memory[nums[i]] {
			memory[nums[i]] = true
			swap(nums, i, j)
			j++
		}
	}
	return j
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	util.PrintNums(nums)
	lenAfter := removeDuplicates(nums)
	fmt.Println(lenAfter)
	util.PrintNums(nums)
}

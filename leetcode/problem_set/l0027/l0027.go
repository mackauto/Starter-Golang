package main

import (
	"fmt"
	"github.com/JPMike/Starter-Golang/leetcode/util"
)

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			swap(nums, i, j)
			j++
		}
	}
	return j
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	util.PrintNums(nums)
	lenAfter := removeElement(nums, val)
	fmt.Println(lenAfter)
	util.PrintNums(nums)
}

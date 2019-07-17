package main

import "github.com/JPMike/Starter-Golang/leetcode/util"

func insertionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j > 0 && nums[j] < nums[j-1]; j-- {
			util.Swap(nums, j-1, j)
		}
	}
}

func main() {
	nums := util.GetExampleNums()
	util.PrintNums(nums)
	insertionSort(nums)
	util.PrintNums(nums)
}

package main

import (
	"github.com/JPMike/Starter-Golang/leetcode/util"
)

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	i := left
	j := right + 1
	for {
		for {
			i++
			if i == right || nums[i] > pivot {
				break
			}
		}
		for {
			j--
			if j == left || nums[j] < pivot {
				break
			}
		}
		if i >= j {
			break
		}
		util.Swap(nums, i, j)
	}
	util.Swap(nums, left, j)
	return j
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// main sort process in partition, which is the unit of work
	index := partition(nums, left, right)
	// recursion in left and right
	quickSort(nums, left, index-1)
	quickSort(nums, index+1, right)
}

func main() {
	nums := util.GetExampleNums()
	util.PrintNums(nums)
	quickSort(nums, 0, len(nums)-1)
	util.PrintNums(nums)
}

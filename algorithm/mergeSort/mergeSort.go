package main

import "github.com/JPMike/Starter-Golang/leetcode/util"

func mergeSort(nums []int) {
	numsLen := len(nums)
	helper := make([]int, numsLen)
	mergeSortDivider(nums, helper, 0, numsLen-1)
}

func mergeSortDivider(nums []int, helper []int, low, high int) {
	if low < high {
		// divide
		middle := (low + high) / 2
		mergeSortDivider(nums, helper, low, middle)
		mergeSortDivider(nums, helper, middle+1, high)
		// merge
		merge(nums, helper, low, middle, high)
	}
}

func merge(nums []int, helper []int, low, middle, high int) {
	// merge two part divided by middle
	// both part are in order
	for i := low; i <= high; i++ {
		helper[i] = nums[i]
	}
	left := low
	right := middle + 1
	cur := low
	// stop if left or right part hit the bound
	for left <= middle && right <= high {
		if helper[left] <= helper[right] {
			nums[cur] = helper[left]
			left++
		} else {
			nums[cur] = helper[right]
			right++
		}
		cur++
	}
	// only left part remain need to be copy back
	// because if right part left, it is in the correct order
	remain := middle - left + 1
	// notice the amount of left remain, how many copy should do
	for i := 0; i < remain; i++ {
		nums[cur+i] = helper[left+i]
	}
}

func main() {
	nums := []int{4, 1, 5, 8, 9, 3, 2, 6, 7}
	util.PrintNums(nums)
	mergeSort(nums)
	util.PrintNums(nums)
}

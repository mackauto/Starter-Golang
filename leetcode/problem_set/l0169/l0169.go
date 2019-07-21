package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	return findMajority(nums, 0, len(nums)-1)
}

func findMajority(nums []int, i, j int) int {
	if i == j {
		return nums[i]
	}
	mid := (i + j) / 2
	leftMajority := findMajority(nums, i, mid)
	rightMajority := findMajority(nums, mid+1, j)
	if leftMajority == rightMajority {
		return leftMajority
	}
	leftMajorityCount := 0
	rightMajorityCount := 0
	for k := i; k <= j; k++ {
		if leftMajority == nums[k] {
			leftMajorityCount += 1
		}
		if rightMajority == nums[k] {
			rightMajorityCount += 1
		}
	}
	if leftMajorityCount > rightMajorityCount {
		return leftMajority
	} else {
		return rightMajority
	}
}

func majorityElementSortMethod(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	type t struct {
		nums []int
		ans  int
	}
	ts := []t{
		{nums: []int{3, 2, 3}, ans: 3},
		{nums: []int{2, 2, 1, 1, 1, 2, 2}, ans: 2},
	}
	for _, t := range ts {
		//ans := majorityElement(t.nums)
		ans := majorityElementSortMethod(t.nums)
		if ans != t.ans {
			fmt.Println(t.nums, "exp:", t.ans, "res:", ans)
			return
		}
	}
	fmt.Println(true)
}

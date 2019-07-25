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
	// divide
	mid := (i + j) / 2
	leftMajority := findMajority(nums, i, mid)
	rightMajority := findMajority(nums, mid+1, j)
	// conquer
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

func majorityElementMapMethod(nums []int) int {
	//memory := map[int]int{}
	memory := make(map[int]int)
	for _, v := range nums {
		memory[v] += 1
		//if _, ok := memory[v]; ok {
		//	memory[v] += 1
		//} else {
		//	memory[v] = 1
		//}
	}
	size := len(nums)
	for k, v := range memory {
		if v > size/2 {
			return k
		}
	}
	return -1
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
		//ans := majorityElementSortMethod(t.nums)
		ans := majorityElementMapMethod(t.nums)
		if ans != t.ans {
			fmt.Println(t.nums, "exp:", t.ans, "res:", ans)
			return
		}
	}
	fmt.Println(true)
}

package main

import (
	"fmt"
	"sort"
)

func smallestRangeI(A []int, K int) int {
	sort.Ints(A)
	lastMin := A[len(A)-1] - K
	firstMax := A[0] + K
	if firstMax >= lastMin {
		return 0
	} else {
		return lastMin - firstMax
	}
}

func main() {
	type t struct {
		nums []int
		k    int
		ans  int
	}
	ts := []t{
		{nums: []int{1}, k: 0, ans: 0},
		{nums: []int{0, 10}, k: 2, ans: 6},
		{nums: []int{1, 3, 6}, k: 3, ans: 0},
	}
	for _, v := range ts {
		if smallestRangeI(v.nums, v.k) != v.ans {
			fmt.Println(false, v)
		}
	}
	fmt.Println(true)
}

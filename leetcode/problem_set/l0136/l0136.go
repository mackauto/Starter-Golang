package main

import "fmt"

func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func main() {
	type t struct {
		nums []int
		ans  int
	}
	ts := []t{
		{nums: []int{2, 2, 1}, ans: 1},
		{nums: []int{4, 1, 2, 1, 2}, ans: 4},
	}
	for _, t := range ts {
		ans := singleNumber(t.nums)
		if ans != t.ans {
			fmt.Println(t.nums, "exp:", t.ans, "res:", ans)
			return
		}
	}
	fmt.Println(true)
}

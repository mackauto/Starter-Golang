package main

import "fmt"

func maxProduct(nums []int) int {
	r := nums[0] // store the result that is the max found so far
	// iMax/iMin stores the max/min product of subarray
	// that ends with the current number nums[i]
	iMax := r
	iMin := r

	for i := 1; i < len(nums); i++ {
		// multiple by a negative makes big number smaller, small number bigger
		if nums[i] < 0 {
			iMax, iMin = iMin, iMax
		}
		iMax = max(nums[i], iMax*nums[i])
		iMin = min(nums[i], iMin*nums[i])
		r = max(r, iMax)
	}
	return r
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func main() {
	type t struct {
		nums []int
		ans  int
	}
	ts := []t{
		{nums: []int{2, 3, -2, 4}, ans: 6},
		{nums: []int{-2, 0, -1}, ans: 0},
	}
	for _, t := range ts {
		ans := maxProduct(t.nums)
		if ans != t.ans {
			fmt.Println(t.nums, "exp:", t.ans, "res:", ans)
			return
		}
	}
	fmt.Println(true)
}

package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	original []int
	nums     []int
}

func Constructor(nums []int) Solution {
	s := Solution{nums: nums, original: make([]int, len(nums))}
	copy(s.original, s.nums)
	return s
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.original)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	for i := range this.nums {
		//rand.Seed(time.Now().Unix())
		swapIndex := rand.Intn(len(this.nums)-i) + i
		this.nums[i], this.nums[swapIndex] = this.nums[swapIndex], this.nums[i]
	}
	return this.nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	obj := Constructor(nums)
	fmt.Println(obj.Shuffle())
	fmt.Println(obj.Reset())
	fmt.Println(obj.Shuffle())
}

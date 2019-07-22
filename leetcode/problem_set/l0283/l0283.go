package main

import "fmt"

func moveZeroes(nums []int) {
	nonZeroIndex := 0
	for cur := 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			swap(nums, nonZeroIndex, cur)
			nonZeroIndex++
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{0, 1, 0, 3, 13}
	moveZeroes(nums)
	fmt.Println(nums)
}

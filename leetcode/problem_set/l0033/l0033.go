package main

import "fmt"

func search(nums []int, target int) int {
	numsLen := len(nums)
	low := 0
	high := numsLen - 1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	rot := low
	low = 0
	high = numsLen - 1
	for low <= high {
		mid := (low + high) / 2
		realMid := (mid + rot) % numsLen
		if nums[realMid] == target {
			return realMid
		}
		if nums[realMid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))
}

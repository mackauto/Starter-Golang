package main

import "fmt"

func isEven(num int) bool {
	return (num & 1) != 1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func sortArrayByParity(A []int) []int {
	i := 0
	j := len(A) - 1
	for i < j {
		// avoid out of index
		for i < len(A) && isEven(A[i]) {
			i++
		}
		for j >= 0 && !isEven(A[j]) {
			j--
		}
		// double check outer loop condition
		if i < j {
			swap(A, i, j)
			i++
			j--
		}
	}
	return A
}

func main() {
	nums := []int{3, 1, 2, 4}
	nums1 := []int{0, 2}
	fmt.Println(sortArrayByParity(nums))
	fmt.Println(sortArrayByParity(nums1))
}

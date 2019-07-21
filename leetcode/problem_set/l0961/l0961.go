package main

import "fmt"

func repeatedNTimes(A []int) int {
	memory := map[int]bool{}
	for i := 0; i < len(A); i++ {
		val := A[i]
		if memory[val] {
			return val
		}
		memory[val] = true
	}
	return -1
}

func repeatedNTimesSubArrayMethod(A []int) int {
	for k := 1; k <= 3; k++ {
		// group start with i in size k
		for i := 0; i < len(A)-k; i++ {
			// compare the first and last element in group
			if A[i] == A[i+k] {
				return A[i]
			}
		}
	}
	return -1
}

func main() {
	nums := []int{5, 1, 5, 2, 5, 3, 5, 4}
	fmt.Println(repeatedNTimes(nums))
	fmt.Println(repeatedNTimesSubArrayMethod(nums))
}

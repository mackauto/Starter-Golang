package main

import "fmt"

func findKthNumber(n int, k int) int {
	cur := 1
	k -= 1 // subtract 1, the cur
	for k > 0 {
		steps := calSteps(n, cur, cur+1)
		if steps <= k {
			// find in the same level
			cur += 1
			k -= steps
		} else {
			// find in the next level, narrow the range
			cur *= 10
			k -= 1 // subtract 1, the up level cur next
		}
	}
	return cur
}

func calSteps(n, cur, next int) int {
	step := 0
	for cur <= n {
		step += min(n+1, next) - cur
		cur *= 10
		next *= 10
	}
	return step
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(findKthNumber(13, 2)) // n=13, k=2, ans=10
}

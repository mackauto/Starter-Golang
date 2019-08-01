package main

import "fmt"

func GetSteps(n, k int) int {
	var dp [100][100]int
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		dp[0][i] = 0
	}
	for j := 1; j <= k; j++ {
		for i := 0; i < n; i++ {
			// ways for point i walk j step to point 0
			dp[j][i] = dp[j-1][(i-1+n)%n] + dp[j-1][(i+1)%n] // for circle, step back is (i-1+n)%n, step forward is (i+1)%n
		}
	}
	return dp[k][0]
}

func main() {
	fmt.Println(GetSteps(10, 2))
}

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	strLen := len(strs)
	if strLen == 0 {
		return ""
	}
	if strLen == 1 {
		return strs[0]
	}
	left := longestCommonPrefix(strs[:strLen/2])
	right := longestCommonPrefix(strs[strLen/2:])
	i := 0
	for i < len(left) && i < len(right) && left[i] == right[i] {
		i++
	}
	return left[:i]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

	strs1 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs1))

	strs2 := []string{"c", "c"}
	fmt.Println(longestCommonPrefix(strs2))
}

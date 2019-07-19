package main

import "fmt"

func perm(a []rune, i int) {
	if i > len(a) {
		fmt.Println(string(a))
		return
	}
	perm(a, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func Perm(a []rune) {
	perm(a, 0)
}

func main() {
	Perm([]rune("abc"))
}

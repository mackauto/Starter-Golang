package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove1(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	data1 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data1))
	fmt.Printf("%q\n", data1)

	s1 := []int{5, 6, 7, 8, 9}
	fmt.Println(remove1(s1, 2))

	s2 := []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(s2, 2))
}

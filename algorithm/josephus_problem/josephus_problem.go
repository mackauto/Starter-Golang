package main

import (
	"container/ring"
	"flag"
	"fmt"
)

func main() {
	total := flag.Int("total", 9, "total")
	num := flag.Int("num", 5, "num")
	fmt.Printf("ring method: %d\n", ringMethod(*total, *num))
	fmt.Printf("recursion method: %d\n", recursionMethod(*total, *num))

	// test case
	// total = 9, num = 5, output = 8
}

func ringMethod(total, num int) interface{} {
	// initial ring 1->2...->9->1...
	r := ring.New(total)
	for i := 1; i < total+1; i++ {
		r.Value = i
		r = r.Next()
	}

	// skip and delete until one left
	for r.Next() != r {
		// find the pre node of node to be deleted
		r = r.Move(num - 2)
		// delete the node
		r.Unlink(1)
		// restart from the next node
		r = r.Next()
	}
	return r.Value
}

func recursionMethod(total, num int) int {
	// base condition
	if total == 1 {
		return total
	}
	// number start from 1
	// old is the number before delete
	// new is the number after delete
	// old = (new + num - 1) % total + 1
	return (recursionMethod(total-1, num)+num-1)%total + 1
}

package main

import "fmt"

// ref: https://draveness.me/golang/basic/golang-interface.html

func Print(v interface{}) {
	// interface{} not represent any type, it is interface{} type during the runtime
	println(v)
}

type TestStruct struct{}

func NilOrNot(v interface{}) {
	if v == nil {
		println("nil")
	} else {
		println("non-nil")
	}
}

func main() {
	type Test struct{}
	v := Test{}
	// type transform when passing v
	Print(v)

	fmt.Println("---")

	var s *TestStruct
	NilOrNot(s)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// use strings.Join for simpler and more efficient
	// any slice may be printed like this
	fmt.Println(strings.Join(os.Args[1:], " "))
}

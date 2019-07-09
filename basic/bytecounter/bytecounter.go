package main

import "fmt"

type ByteCounter int

// receiver type is pointer
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	_, _ = c.Write([]byte("hello"))
	fmt.Println(c)
	// call Write func of c, implement of io.Writer interface
	_, _ = fmt.Fprintf(&c, "hello")
	fmt.Println(c)

	// output:
	// 5
	// 10
}

package command_line

import (
	"fmt"
	"os"
	"testing"
)

// echo command line arguments
func Echo1() {
	var s, sep string
	// how to get argument from command line
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func TestEcho1(t *testing.T) {
	Echo1()
}

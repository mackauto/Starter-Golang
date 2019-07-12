package main

import "fmt"

func main() {
	// should avoid pointer chan, or notice for range
	ch := make(chan *int, 3)
	input := []int{1, 2, 3}
	go func() {
		for _, v := range input {
			//only send the address of variable v, which is not we wanted
			//ch <- &v

			// fix: use local variable temp
			temp := v
			ch <- &temp
		}
		// fix: you can fix that this way
		for k := range input {
			ch <- &input[k]
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(*v)
	}
	// output: 3 3 3 1 2 3
}

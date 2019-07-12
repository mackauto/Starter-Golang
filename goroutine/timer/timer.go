package main

import (
	"fmt"
	"time"
)

func timer(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}

func main() {
	start := time.Now().Unix()
	for i := 0; i < 10; i++ {
		c := timer(time.Duration(i) * time.Second)
		<-c
	}
	end := time.Now().Unix()
	fmt.Println(end - start)
	// output should be (0+9)*10/2=45
}

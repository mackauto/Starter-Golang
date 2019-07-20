package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				// let other go routine to run, but do not suspend the current routine
				runtime.Gosched()
			}
		}()
	}
	// let go routine run for 1 sec
	time.Sleep(time.Second)

	// copy counter value for safe
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", opsFinal)
}

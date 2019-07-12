package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	WORKERS    = 5
	SUBWORKERS = 3
	TASKS      = 20
	SUBTASKS   = 10
)

func subWorker(subTasks chan int) {
	for {
		task, ok := <-subTasks
		if !ok {
			return
		}
		time.Sleep(time.Duration(task) * time.Millisecond)
		fmt.Println(task)
	}
}

func worker(tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			return
		}

		subTasks := make(chan int)

		for i := 0; i < SUBWORKERS; i++ {
			go subWorker(subTasks)
		}

		for i := 0; i < SUBTASKS; i++ {
			task1 := task * i
			subTasks <- task1
		}
		close(subTasks)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(WORKERS)
	tasks := make(chan int)

	// create worker first, fetch first for channel without buffer
	// fetch last for channel with buffer
	for i := 0; i < WORKERS; i++ {
		go worker(tasks, &wg)
	}

	for i := 0; i < TASKS; i++ {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
}

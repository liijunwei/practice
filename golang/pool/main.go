package main

import (
	"log"
	"time"
)

func main() {
	startTime := time.Now()

	count := 20
	tasks := make([]Task, count)

	for i := 0; i < count; i++ {
		tasks[i] = Task{ID: i + 1}
	}

	concurrency := 8
	pool := NewPool(tasks, concurrency)
	pool.Run()

	log.Println("all done, elapsed: ", time.Since(startTime))
}

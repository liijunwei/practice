package main

import (
	"log"
	"time"
)

func main() {
	startTime := time.Now()

	tasks := []Task{
		NewProcessImageTask(8),
		NewProcessImageTask(2),
		NewSendEmailTask(3),
		NewSendEmailTask(4),
		NewProcessImageTask(4),
		NewProcessImageTask(3),
		NewProcessImageTask(6),
		NewSendEmailTask(1),
		NewSendEmailTask(5),
		NewSendEmailTask(2),
		NewProcessImageTask(1),
		NewSendEmailTask(7),
		NewSendEmailTask(6),
		NewProcessImageTask(7),
		NewProcessImageTask(5),
		NewSendEmailTask(8),
	}

	concurrency := 10
	pool := NewPool(tasks, concurrency)
	pool.Run()

	log.Println("all done, elapsed: ", time.Since(startTime))
}

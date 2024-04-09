package main

import (
	"log"
	"time"
)

func main() {
	startTime := time.Now()

	concurrency := 5
	tasks := make(chan int)
	results := make(chan int)

	for i := 0; i < concurrency; i++ {
		go worker(i+1, tasks, results)
	}

	jobCount := 20
	for i := 0; i < jobCount; i++ {
		tasks <- i + 1
	}

	close(tasks)

	for {
		select {
		case _, ok := <-results:
			if !ok {
				break
			}
		default:
		}
	}

	log.Println("all done, elapsed: ", time.Since(startTime))
}

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		log.Printf("worker#%d processing... %d", id, task)
		time.Sleep(1 * time.Second)
		results <- task * 2
	}
}

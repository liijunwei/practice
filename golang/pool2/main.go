package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	concurrency := 5
	jobCount := 10
	tasks := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for id := range concurrency {
		go func() {
			worker(id+1, tasks, results)
			wg.Done()
		}()
	}

	go func() {
		for i := 0; i < jobCount; i++ {
			tasks <- i + 1
		}
		close(tasks)

		wg.Wait()
		close(results)
	}()

	for r := range results {
		log.Println("result", r)
	}

	log.Println("all done, elapsed: ", time.Since(startTime))
}

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		log.Printf("worker#%d processing... %d", id, task)
		time.Sleep(100 * time.Millisecond)
		results <- task * 2
	}
}

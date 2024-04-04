package main

import (
	"log"
	"sync"
	"time"
)

type Task struct {
	ID int
}

func (t *Task) Run() {
	log.Printf("processing task %d\n", t.ID)
	time.Sleep(2 * time.Second)
}

type Pool struct {
	Tasks       []Task
	concurrency int
	taskChannel chan Task
	wg          sync.WaitGroup
}

func NewPool(tasks []Task, concurrency int) Pool {
	return Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		taskChannel: make(chan Task),
	}
}

func (p *Pool) startWorker() {
	for task := range p.taskChannel {
		p.wg.Add(1)
		task.Run()
		p.wg.Done()
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.startWorker()
	}

	for _, task := range p.Tasks {
		p.taskChannel <- task
	}

	close(p.taskChannel)

	p.wg.Wait()
}

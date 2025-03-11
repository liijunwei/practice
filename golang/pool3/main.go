package main

import (
	"sync"
	"time"
)

func main() {
	tasks := []Task{
		&t1{},
		&t2{},
		&t1{},
		&t2{},
		&t1{},
		&t2{},
	}

	start := time.Now()

	pool := NewPool(tasks, 100)
	pool.Run()
	println("all done, elapsed:", time.Since(start).String())
}

type Task interface {
	Run()
}

type Pool struct {
	Tasks       []Task
	concurrency int
	taskChannel chan Task
	wg          sync.WaitGroup
}

func NewPool(tasks []Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		taskChannel: make(chan Task),
	}
}

func (p *Pool) Run() {
	for range p.concurrency {
		go p.startWorker()
	}

	p.wg.Add(len(p.Tasks))

	for _, task := range p.Tasks {
		p.taskChannel <- task
	}

	p.wg.Wait()
	close(p.taskChannel)
}

func (p *Pool) startWorker() {
	for t := range p.taskChannel {
		p.runOne(t)
	}
}

func (p *Pool) runOne(task Task) {
	defer p.wg.Done()

	task.Run()
}

type t1 struct{}

func (t *t1) Run() {
	println("working on t1")
	time.Sleep(100 * time.Millisecond)
}

type t2 struct{}

func (t *t2) Run() {
	println("working on t2")
	time.Sleep(10 * time.Millisecond)
}

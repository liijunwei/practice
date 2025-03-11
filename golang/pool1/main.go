package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	tasks := []Task{
		NewProcessImageTask(1),
		NewProcessImageTask(2),
		NewProcessImageTask(3),
		NewProcessImageTask(4),
		NewProcessImageTask(5),
		NewProcessImageTask(6),
		NewProcessImageTask(7),
		NewProcessImageTask(8),
		NewSendEmailTask(1),
		NewSendEmailTask(2),
		NewSendEmailTask(3),
		NewSendEmailTask(4),
		NewSendEmailTask(5),
		NewSendEmailTask(6),
		NewSendEmailTask(7),
		NewSendEmailTask(8),
	}

	concurrency := 10
	pool := NewPool(tasks, concurrency)
	pool.Run()

	log.Println("all done, elapsed: ", time.Since(startTime))
}

type Task interface {
	Run()
}

type SendEmail struct {
	Email       string
	Subject     string
	MessageBody string
}

func NewSendEmailTask(id int) *SendEmail {
	return &SendEmail{
		Email:       fmt.Sprintf("me-%d@foo.com", id),
		Subject:     fmt.Sprintf("fiz-%d", id),
		MessageBody: fmt.Sprintf("buz-%d", id),
	}
}

func (t *SendEmail) Run() {
	content, err := json.Marshal(t)
	boom(err)

	log.Printf("processing... %s\n", content)
	time.Sleep(2 * time.Second)
}

func NewProcessImageTask(id int) *ProcessImage {
	return &ProcessImage{
		ImageUrl: fmt.Sprintf("image@%d", id),
	}
}

type ProcessImage struct {
	ImageUrl string
}

func (t *ProcessImage) Run() {
	content, err := json.Marshal(t)
	boom(err)

	log.Printf("processing... %s\n", content)
	time.Sleep(5 * time.Second)
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

func boom(err error) {
	if err != nil {
		panic(err)
	}
}

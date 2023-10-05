package main

import (
	"log"
	"os"
	"time"
)

var query = "test"
var matches int

var workerCount = 0
var maxWorkerCount = 32
var searchRequest = make(chan string)
var workDone = make(chan bool)
var foundMatch = make(chan bool)

func main() {
	start := time.Now()

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	startPoint := dirname + "/"
	workerCount = 1
	go search(startPoint, true)

	wait()

	log.Println(matches, "matches")

	log.Println(time.Since(start))
}

func wait() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path, true)
		case <-workDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			matches++
		}
	}
}

func search(path string, master bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Println("cannot read dir: ", path)
		workDone <- true

		return
	}

	for _, file := range files {
		name := file.Name()

		if name == query {
			foundMatch <- true
		}

		if file.IsDir() {
			newDir := path + name + "/"

			if workerCount < maxWorkerCount {
				searchRequest <- newDir
			} else {
				search(newDir, false)
			}
		}
	}

	if master {
		workDone <- true
	}
}

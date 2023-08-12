package main

import (
	"fmt"
	"time"
)

const (
	statusHold     = "hold"
	statusReady    = "ready"
	statusRunning  = "running"
	statusFinished = "finished"
)

func main() {
	var taskGraph = map[string][]string{
		"t2": {"t1"},
		"t3": {"t1"},
		"t4": {"t3"},
	}

	fmt.Println(taskGraph)
}

type WorkNode struct {
	Name     string
	Duration time.Duration
	Status   string
}

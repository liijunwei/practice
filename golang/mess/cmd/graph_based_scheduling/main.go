package main

import (
	"fmt"
	"time"

	"github.com/dominikbraun/graph"
)

const (
	statusHold     = "hold"
	statusReady    = "ready"
	statusRunning  = "running"
	statusFinished = "finished"
)

func main() {
	fmt.Println("graph based scheduling demo")

	nodeHash := func(c WorkNode) string {
		return c.Name
	}

	g1 := graph.New(nodeHash, graph.Directed(), graph.PreventCycles())
	n1 := WorkNode{
		Name:           "task1",
		DurationNeeded: 10 * time.Millisecond,
		Status:         "hold",
	}

	n2 := WorkNode{
		Name:           "task2",
		DurationNeeded: 20 * time.Millisecond,
		Status:         "hold",
	}

	n3 := WorkNode{
		Name:           "task3",
		DurationNeeded: 30 * time.Millisecond,
		Status:         "hold",
	}

	n4 := WorkNode{
		Name:           "task4",
		DurationNeeded: 40 * time.Millisecond,
		Status:         "hold",
	}

	_ = g1.AddVertex(n1)
	_ = g1.AddVertex(n2)
	_ = g1.AddVertex(n3)
	_ = g1.AddVertex(n4)

	_ = g1.AddEdge(n1.Name, n2.Name)
	_ = g1.AddEdge(n1.Name, n3.Name)
	_ = g1.AddEdge(n3.Name, n4.Name)

	order, _ := graph.TopologicalSort(g1)
	fmt.Println(order)
}

type WorkNode struct {
	Name           string
	DurationNeeded time.Duration
	Status         string
}

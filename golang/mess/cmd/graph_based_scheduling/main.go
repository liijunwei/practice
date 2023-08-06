package main

import (
	"fmt"
	"time"

	"github.com/dominikbraun/graph"
)

func main() {
	fmt.Println("graph based scheduling demo")

	nodeHash := func(c WorkNode) string {
		return c.Name
	}

	g1 := graph.New(nodeHash, graph.Directed(), graph.PreventCycles())
	n1 := WorkNode{
		Name:              "task1",
		MillisecondNeeded: 10,
	}

	n2 := WorkNode{
		Name:              "task2",
		MillisecondNeeded: 20,
	}

	n3 := WorkNode{
		Name:              "task3",
		MillisecondNeeded: 30,
	}

	n4 := WorkNode{
		Name:              "task4",
		MillisecondNeeded: 40,
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
	Name              string
	MillisecondNeeded time.Duration
}

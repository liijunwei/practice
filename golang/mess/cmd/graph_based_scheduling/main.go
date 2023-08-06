package main

import (
	"fmt"
	"sync"
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

	// adjMap, err := g1.AdjacencyMap()
	// if err != nil {
	// 	panic(err)
	// }

	for _, node := range order {
		// fmt.Println(adjMap[node])
		fmt.Printf("working on %s \n", node)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	queue := make(chan string)

	go func() {
		produce(queue, g1)
		wg.Done()
	}()

	go func() {
		consume(queue, g1)
		wg.Done()
	}()

	wg.Wait()
}

type WorkNode struct {
	Name           string
	DurationNeeded time.Duration
	Status         string
}

func produce(queue chan<- string, _graph graph.Graph[string, WorkNode]) {
	for i := 0; i < 10; i++ {
		queue <- fmt.Sprintf("task#%d", i)
	}

	close(queue)
}

func consume(queue <-chan string, _graph graph.Graph[string, WorkNode]) {
	for {
		select {
		case task, ok := <-queue:
			if ok {
				fmt.Println("working on ", task)
			} else {
				fmt.Println("queue closed")
				return
			}
		default:
			fmt.Println("idle...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

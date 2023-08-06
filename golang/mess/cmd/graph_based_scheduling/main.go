package main

import (
	"fmt"

	"github.com/dominikbraun/graph"
)

func main() {
	fmt.Println("graph based scheduling demo")

	g := graph.New(graph.IntHash, graph.Directed())

	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_ = g.AddVertex(3)
	_ = g.AddVertex(4)

	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(1, 3)
	_ = g.AddEdge(3, 4)

	_ = graph.DFS(g, 1, func(value int) bool {
		fmt.Println(value)
		return false
	})
}

type WorkNode struct {
}

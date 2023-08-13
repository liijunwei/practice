package dfs

import (
	"fmt"
	"testing"
)

type Graph struct {
	adjList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[string][]string),
	}
}

func (g *Graph) AddEdge(u, v string) {
	g.adjList[u] = append(g.adjList[u], v)
}

func DFS(g *Graph, start string, visited map[string]bool) {
	visited[start] = true
	fmt.Printf("visiting %s\n", start)

	neighbors := g.adjList[start]
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			DFS(g, neighbor, visited)
		}
	}
}

func Test_DFS(t *testing.T) {
	g := NewGraph()

	g.AddEdge("A", "C")
	g.AddEdge("A", "D")
	g.AddEdge("B", "C")
	g.AddEdge("B", "D")
	g.AddEdge("D", "E")
	g.AddEdge("D", "F")
	g.AddEdge("C", "G")
	g.AddEdge("C", "H")
	g.AddEdge("E", "H")
	g.AddEdge("F", "I")
	g.AddEdge("G", "J")
	g.AddEdge("I", "J")
	g.AddEdge("H", "J")
	g.AddEdge("J", "K")
	g.AddEdge("J", "L")
	g.AddEdge("J", "M")

	startNode := "A"
	// startNode := "B"
	visited := make(map[string]bool)

	fmt.Println("Depth-First Traversal (DFS):")
	DFS(g, startNode, visited)
}

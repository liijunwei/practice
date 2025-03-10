package main

// https://github.com/adonovan/gopl.io/blob/master/ch5/toposort/main.go
// See page 136.

// The toposort program prints the nodes of a DAG in topological order.

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%03d: %s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	// Define DFS function that takes graph and current vertex
	var dfs func(g map[string][]string, v string)

	dfs = func(g map[string][]string, v string) {
		if !seen[v] {
			seen[v] = true
			// Visit all prerequisites (adjacent vertices)
			for _, prereq := range g[v] {
				dfs(g, prereq)
			}
			order = append(order, v)
		}
	}

	// Get all courses (vertices)
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Call DFS on each course
	for _, key := range keys {
		dfs(m, key)
	}

	return order
}

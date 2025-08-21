package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func dumpdata(lst any) string {
	data, err := json.Marshal(lst)
	boom(err)
	return string(data)
}

func dumpgraph(heap *MaxHeap, dotfilename string) {
	var content strings.Builder
	content.WriteString("digraph heap {\n")
	content.WriteString("    node [shape=circle];\n")

	for i, val := range heap.pq {
		content.WriteString(fmt.Sprintf("    node%d [label=\"%d\"];\n", i, val))
	}

	for i := range heap.pq {
		if idx := left(i); idx < len(heap.pq) {
			content.WriteString(fmt.Sprintf("    node%d -> node%d;\n", i, idx))
		}

		if idx := right(i); idx < len(heap.pq) {
			content.WriteString(fmt.Sprintf("    node%d -> node%d;\n", i, idx))
		}
	}
	content.WriteString("}\n")

	svgfilename := strings.ReplaceAll(dotfilename, ".dot", ".svg")
	boom(os.WriteFile(dotfilename, []byte(content.String()), 0644))
	boom(exec.Command("dot", "-Tsvg", dotfilename, "-o", svgfilename).Run())
	boom(exec.Command("open", svgfilename).Run())
	println(dotfilename)
	println(svgfilename)
}

func boom(err error) {
	if err != nil {
		panic(err)
	}
}

func assert(ok bool) {
	if !ok {
		panic("ASSERTION FAILED")
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	h := NewMaxHeap()
	const N = 20
	lst := make([]int, 0, N)
	deleted := make([]int, 0, N)
	var seed int64 = 2026
	gen := rand.New(rand.NewSource(seed))

	for range N {
		num := gen.Intn(1000)
		lst = append(lst, num)
		h.Insert(num)
	}
	fmt.Println(h)
	h.Dump("/tmp/max_heap.dot")

	for range N {
		deleted = append(deleted, h.DeleteMax())
	}
	sort.Slice(lst, func(i, j int) bool {
		return lst[i] > lst[j]
	})
	assert(equal(lst, deleted))
	fmt.Println(h)
	fmt.Println("pass")
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		array: make([]int, 0, 100),
	}
}

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	lastElementIdx := len(h.array) - 1
	h.swim(lastElementIdx)
}

func (h *MaxHeap) DeleteMax() int {
	max := h.array[0]
	lastElementIdx := len(h.array) - 1
	if lastElementIdx < 0 {
		return lastElementIdx
	}
	h.array[0] = h.array[lastElementIdx]
	h.array = h.array[:lastElementIdx]
	h.sink(0)
	return max
}

func (h *MaxHeap) swim(index int) {
	assert(index >= 0)
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) sink(index int) {
	l, r, last := left(index), right(index), h.last()

	tocompare := 0

	for l <= last {
		if l == last {
			tocompare = l
		} else if h.array[l] > h.array[r] {
			tocompare = l
		} else {
			tocompare = r
		}

		if h.array[index] >= h.array[tocompare] {
			return
		}

		h.swap(index, tocompare)
		index = tocompare
		l, r = left(index), right(index)
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

// last element index
func (h *MaxHeap) last() int {
	return len(h.array) - 1
}

func (h *MaxHeap) Dump(dotfilename string) {
	var content strings.Builder
	content.WriteString("digraph heap {\n")
	content.WriteString("    node [shape=circle];\n")

	for i, val := range h.array {
		content.WriteString(fmt.Sprintf("    node%d [label=\"%d\"];\n", i, val))
	}

	for i := range h.array {
		if idx := left(i); idx < len(h.array) {
			content.WriteString(fmt.Sprintf("    node%d -> node%d;\n", i, idx))
		}

		if idx := right(i); idx < len(h.array) {
			content.WriteString(fmt.Sprintf("    node%d -> node%d;\n", i, idx))
		}
	}
	content.WriteString("}\n")

	svgfilename := strings.ReplaceAll(dotfilename, ".dot", ".svg")
	boom(os.WriteFile(dotfilename, []byte(content.String()), 0644))
	boom(exec.Command("dot", "-Tsvg", dotfilename, "-o", svgfilename).Run())
	boom(exec.Command("open", svgfilename).Run())
	fmt.Println(dotfilename)
	fmt.Println(svgfilename)
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

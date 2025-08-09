package main

import (
	"fmt"
	"math/rand/v2"
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
	for range N {
		num := rand.IntN(1000)
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

func (h *MaxHeap) Dump(filename string) {
	content := "digraph heap {\n"
	content += "    node [shape=circle];\n"

	for i, val := range h.array {
		content += fmt.Sprintf("    node%d [label=\"%d\"];\n", i, val)
	}

	for i := 0; i < len(h.array); i++ {
		leftIdx := left(i)
		if leftIdx < len(h.array) {
			content += fmt.Sprintf("    node%d -> node%d;\n", i, leftIdx)
		}

		rightIdx := right(i)
		if rightIdx < len(h.array) {
			content += fmt.Sprintf("    node%d -> node%d;\n", i, rightIdx)
		}
	}

	content += "}\n"
	boom(os.WriteFile(filename, []byte(content), 0644))

	svgfilename := strings.ReplaceAll(filename, ".dot", ".svg")
	err := exec.Command("dot", "-Tsvg", filename, "-o", svgfilename).Run()
	boom(err)
	err = exec.Command("open", svgfilename).Run()
	boom(err)
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

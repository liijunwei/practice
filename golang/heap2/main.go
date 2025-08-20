package main

import (
	"encoding/json"
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
	h.Dump("/tmp/max_heap.dot")

	for range N {
		deleted = append(deleted, h.DeleteMax())
	}
	println("lst raw   ", dump(lst))
	sort.Slice(lst, func(i, j int) bool {
		return lst[i] > lst[j]
	})
	println("lst sorted", dump(lst))
	println("deleted   ", dump(deleted))
	assert(equal(lst, deleted))
	println("pass")
}

func dump(lst any) string {
	data, err := json.Marshal(lst)
	boom(err)
	return string(data)
}

func NewMaxHeap() *MaxHeap {
	arr := []int{-1} //ignore first element
	return &MaxHeap{
		array: arr,
	}
}

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.swim(h.N())
	// h.Dump(fmt.Sprintf("fooo-%d.dot", key))
}

func (h *MaxHeap) DeleteMax() int {
	max := h.array[1] //ignore the first array element
	h.swap(1, h.N())
	h.array = h.array[:h.len()-1]
	h.sink(1)
	return max
}

func (h *MaxHeap) swim(k int) {
	assert(k >= 0 && k <= h.len())
	for k > 1 && h.less(parent(k), k) {
		h.swap(parent(k), k)
		k = parent(k)
	}
}

// 下沉操作, 将元素向下移动以维护堆的性质
func (h *MaxHeap) sink(k int) {
	for left(k) <= h.N() {
		// 左子节点
		j := left(k)
		// 选择两个子节点中中较大的一个
		if j < h.N() && h.less(j, j+1) {
			j++
		}
		// 如果当前节点大于等于子节点, 则可以停止sink
		if !h.less(k, j) {
			break
		}
		// 交换
		h.swap(k, j)
		// 继续迭代
		k = j
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *MaxHeap) len() int {
	assert(len(h.array) >= 1)
	assert(len(h.array) == h.N()+1)
	return len(h.array)
}

func (h *MaxHeap) N() int {
	return len(h.array) - 1
}

func (h *MaxHeap) less(i, j int) bool {
	return h.array[i] < h.array[j]
}

//		 k/2
//		  |
//			k
//	 /    \
//
// 2k      2k+1
func parent(k int) int {
	return k / 2
}

func left(k int) int {
	return 2 * k
}

func right(k int) int {
	return 2*k + 1
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

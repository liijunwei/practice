package main

import (
	"fmt"
	"math"
)

// TODO use quickcheck test
// TODO build heap from top to bottom
// TODO build heap from bottom to top
// TODO compare the efficiency
func main() {
	lst := []int{888, 403, 907, 349, 639, 372, 371, 409, 127, 735, 137, 896, 185, 361, 177, 703, 806, 393, 786, 148}
	h := NewMaxHeap(len(lst))
	for _, n := range lst {
		h.Insert(n)
	}
}

type MaxHeap struct {
	pq   []int
	size int
}

func NewMaxHeap(maxN int) *MaxHeap {
	return &MaxHeap{
		pq: make([]int, maxN+1),
	}
}

func (h *MaxHeap) Insert(v int) {
	h.size++
	h.pq[h.size] = v
	h.swim(h.size)
	dumpgraph(h, fmt.Sprintf("tmp/heap-%d.dot", h.size))
}

func (h *MaxHeap) DelMax() int {
	max := h.pq[1]
	h.exch(1, h.size)
	h.pq[h.size+1] = int(math.Inf(-1))
	h.size--
	h.sink(1)
	return max
}

func (h *MaxHeap) Size() int {
	return h.size
}

func (h *MaxHeap) swim(k int) {
	for k > 1 && h.less(parent(k), k) {
		h.exch(parent(k), k)
		k = parent(k)
	}
}

func (h *MaxHeap) sink(k int) {
	for left(k) <= h.size {
		// left child
		j := left(k)

		// find bigger child
		if j < h.size && h.less(j, j+1) {
			j++
		}

		// 当前节点大于等于任意子节点时停止
		if !h.less(k, j) {
			break
		}

		// 当前节点小余子节点时交换，并继续下次迭代
		h.exch(k, j)
		k = j
	}
}

func (h *MaxHeap) less(i, j int) bool {
	return h.pq[i] < h.pq[j]
}

func (h *MaxHeap) exch(i, j int) {
	h.pq[i], h.pq[j] = h.pq[j], h.pq[i]
}

func parent(k int) int {
	return k / 2
}

func left(k int) int {
	return 2 * k
}

func right(k int) int {
	return 2*k + 1
}

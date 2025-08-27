package main

import "fmt"

func main() {
	lst := []int{888, 403, 907, 349, 639, 372, 371, 409, 127, 735, 137, 896, 185, 361, 177, 703, 806, 393, 786, 148}
	h := NewMaxHeap(len(lst))
	for _, n := range lst {
		h.Insert(n)
	}

	fmt.Println(h)
}

type MaxHeap struct {
	pq   []int
	size int
}

func NewMaxHeap(maxN int) *MaxHeap {
	return &MaxHeap{
		pq:   make([]int, maxN+1),
		size: 0,
	}
}

func (h *MaxHeap) Insert(v int) {
	h.size++
	h.pq[h.size] = v
	h.swim(h.size)
}

func (h *MaxHeap) DelMax() int {
	max := h.pq[1] //ignore first array element
	h.exch(1, h.size)
	h.size--
	h.sink(1)
	return max
}

func (h *MaxHeap) Size() int {
	return h.size
}

func (h *MaxHeap) Empty() bool {
	return h.Size() == 0
}

func (h *MaxHeap) swim(k int) {
	for k > 1 && h.less(parent(k), k) {
		h.exch(parent(k), k)
		k = parent(k)
	}
}

func (h *MaxHeap) sink(k int) {
	for left(k) <= h.size {
		//left child
		j := left(k)

		//find the bigger child
		if j < h.size && h.less(j, j+1) {
			j++
		}

		// break the look when heap property is satisfied
		if !h.less(k, j) {
			break
		}

		// exch and continue sink when heap property is not satisfied
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
	return k * 2
}

func right(k int) int {
	return k*2 + 1
}

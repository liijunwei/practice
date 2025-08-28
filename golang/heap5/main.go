package main

import "fmt"

func main() {
	h := NewMaxHeap(10)
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	fmt.Println("pq", h.pq)
	fmt.Println("delmax", h.DelMax(), h.DelMax(), h.DelMax())
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

func (h *MaxHeap) Insert(val int) {
	h.size++
	h.pq[h.size] = val
	h.up(h.size)
}

func (h *MaxHeap) DelMax() int {
	max := h.pq[1]
	h.swap(1, h.size)
	h.size--
	h.down(1)
	return max
}

func (h *MaxHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *MaxHeap) up(k int) {
	for k > 1 && h.less(k/2, k) {
		h.swap(k/2, k)
		k = k / 2
	}
}

func (h *MaxHeap) down(k int) {
	for 2*k <= h.size {
		j := 2 * k                        // left child
		if j < h.size && h.less(j, j+1) { // left child < right child
			j++
		}

		if !h.less(k, j) { // current >= bigger child
			break
		}

		// current < bigger child
		h.swap(k, j)

		// next iteration
		k = j
	}
}

func (h *MaxHeap) swap(j, k int) {
	h.pq[j], h.pq[k] = h.pq[k], h.pq[j]
}

func (h *MaxHeap) less(j, k int) bool {
	return h.pq[j] < h.pq[k]
}

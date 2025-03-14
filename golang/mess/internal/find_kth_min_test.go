package internal

import (
	"fmt"
	"testing"

	"github.com/tidwall/tinyqueue"
)

// 从两个有序数组的并集中寻找第k小元素

// http://bookshadow.com/weblog/2015/02/05/find-kth-smallest-element-in-the-union-of-two-sorted-arrays/

func TestFindKthMin(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5, 6}
	a2 := []int{6, 7, 8, 9, 10, 11, 12}

	k := 8
	bruteforce(a1, a2, k)
	fmt.Println("----")
	pqFind(a1, a2, k)
	fmt.Println("----")
	v3Find(a1, a2, k)
}

type intval int

func (a intval) Less(b tinyqueue.Item) bool {
	return a < b.(intval)
}

func pqFind(a1, a2 []int, k int) {
	pq := tinyqueue.New(nil)

	for _, e := range a1 {
		pq.Push(intval(e))
	}

	for _, e := range a2 {
		pq.Push(intval(e))
	}

	var result tinyqueue.Item

	for range k {
		result = pq.Pop()
	}

	fmt.Printf("%dth min is %d\n", k, result)
}

func bruteforce(a1, a2 []int, k int) {
	r := merge(a1, a2)
	fmt.Println(r)

	fmt.Printf("%dth min is %d\n", k, r[k-1])
}

func merge(a1, a2 []int) []int {
	m := len(a1)
	n := len(a2)
	r := make([]int, m+n)
	copy(r, a1)
	a1 = r

	idx := m + n - 1

	i := m - 1
	j := n - 1

	for i >= 0 && j >= 0 {
		if a1[i] >= a2[j] {
			a1[idx] = a1[i]
			i--
		} else {
			a1[idx] = a2[j]
			j--
		}
		idx--
	}

	for j >= 0 {
		a1[idx] = a2[j]
		idx--
		j--
	}

	return a1
}

func v3Find(A, B []int, k int) {
	result := okSolution(A, B, k)
	fmt.Printf("%dth min is %d\n", k, result)
}

// unclear yet
func okSolution(A, B []int, k int) int {
	a_offset := 0
	b_offset := 0

	if len(A)+len(B) < k {
		panic("should not happen")
	}

	for {
		if a_offset < len(A) {
			for b_offset == len(B) || (a_offset < len(A) && A[a_offset] <= B[b_offset]) {
				fmt.Println("a_offset", a_offset)
				a_offset++

				if a_offset+b_offset == k {
					return A[a_offset]
				}
			}
		}

		if b_offset < len(B) {
			for a_offset == len(A) || A[a_offset] >= B[b_offset] {
				b_offset++

				if a_offset+b_offset == k {
					return B[b_offset]
				}
			}
		}
	}
}

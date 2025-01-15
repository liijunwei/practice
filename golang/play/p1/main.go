package main

import (
	"cmp"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	t1()
	// t2()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(n int) *ListNode {
	newNode := new(ListNode)
	newNode.Val = n
	l.Next = newNode

	return newNode
}

func t1() {
	l1 := &ListNode{Val: 9}
	l1.add(9).add(9).add(9).add(9).add(9).add(9)

	l2 := &ListNode{Val: 9}
	l2.add(9).add(9).add(9)

	demo := addTwoNumbers(l1, l2)
	spew.Dump(demo)
}

func t2() {
	l1 := &ListNode{Val: 2}
	l1.add(4).add(3)

	l2 := &ListNode{Val: 5}
	l2.add(6).add(4)

	demo := addTwoNumbers(l1, l2)
	spew.Dump(demo)
}

// 342+465=807
// 0+0=0
// 9999999+9999=10009998
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	counter := 0

	head := new(ListNode)
	curr := head
	t1 := l1
	t2 := l2
	carry := 0

	for t1 != nil || t2 != nil {
		counter += 1
		if counter > 1000 {
			panic("maybe dead loop")
		}

		println("carry", carry)

		a := cmp.Or(t1, &ListNode{}).Val
		b := cmp.Or(t2, &ListNode{}).Val

		val := a + b + carry
		carry = val / 10
		val = val % 10

		curr.Next = &ListNode{Val: val}
		curr = curr.Next

		if t1 != nil {
			t1 = t1.Next
		}

		if t2 != nil {
			t2 = t2.Next
		}
	}

	if carry == 1 {
		curr.Next = &ListNode{Val: carry}
	}

	return head.Next
}

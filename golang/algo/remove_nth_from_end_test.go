package algo

import (
	"fmt"
	"testing"
)

func TestFindNthFromEnd(t *testing.T) {
	printLinkedList(newList())
	fmt.Println()
	for i := range 5 {
		printLinkedList(removeNthFromEnd1(i+1, newList()))
	}
	fmt.Println()
	for i := range 5 {
		printLinkedList(removeNthFromEnd2(i+1, newList()))
	}
}

func newList() *Node {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	return n1
}

/**
 * Definition for singly-linked list.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */
func removeNthFromEnd1(n int, head *Node) *Node {
	if head == nil {
		return nil
	}

	prev := head
	curr := head

	// assume n < list length
	for range n {
		curr = curr.Next
	}

	if curr == nil {
		return head.Next
	}

	for curr.Next != nil {
		prev = prev.Next
		curr = curr.Next
	}

	prev.Next = prev.Next.Next

	return head
}

type Node struct {
	Val  int
	Next *Node
}

func printLinkedList(head *Node) {
	curr := head

	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

// this is smart
func removeNthFromEnd2(n int, head *Node) *Node {
	dummy := &Node{Next: head}
	left, right := dummy, dummy

	for range n {
		right = right.Next
	}

	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return dummy.Next
}

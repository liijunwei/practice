package main

import (
	"fmt"
	"os"
	"strconv"
)

// linked list, delete last K th node
//
// e.g.
// 1,2,3,4,5 1 => 1,2,3,4
// 1,2,3,4,5 2 => 1,2,3,5
// 1,2,3,4,5 3 => 1,2,4,5
// 1,2,3,4,5 4 => 1,3,4,5
// 1,2,3,4,5 5 => 2,3,4,5
// 1,2,3,4,5 6 => 2,3,4,5

// go run linkedlist/lastkth/main.go 1
// go run linkedlist/lastkth/main.go 2

// use two pointers and maintain a fixed length window in between(sliding window?)
//
// no need to find the list length
// first move the pointer `tail` k times and then use another pointer `prev` to move along with `tail`
// the distance between `prev` and `tail` is k
func main() {
	head := BuildLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Print("init: ")
	PrintLinkedList(head)

	k, _ := strconv.Atoi(os.Args[1])

	newHead := solution(head, k)
	fmt.Print("old_head: ")
	PrintLinkedList(head)
	fmt.Print("new_head: ")
	PrintLinkedList(newHead)
}

func solution(head *Node, k int) *Node {
	if head == nil || k <= 0 {
		fmt.Println("empty list or invalid k value")
		return nil
	}

	tail := head

	// Q: how to handle the case when k > list length?
	// A: choose fallback k to the list length
	for range k {
		tail = tail.Next

		if tail == nil {
			break
		}
	}

	prev := head

	if tail == nil {
		prev = prev.Next
		fmt.Println("k", k, "prev", prev.Val, "tail", tail)
		head.Next = nil

		return prev
	}

	for tail.Next != nil {
		tail = tail.Next
		prev = prev.Next
	}

	// fmt.Println("k", k, "prev", prev.Val, "tail", tail.Val)
	prev.Next = prev.Next.Next

	return head
}

type Node struct {
	Val  int
	Next *Node
}

func BuildLinkedList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	head := &Node{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		newNode := &Node{Val: val}
		current.Next = newNode
		current = newNode
	}
	return head
}

func PrintLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

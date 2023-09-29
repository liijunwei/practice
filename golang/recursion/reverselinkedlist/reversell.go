package reversell

import "fmt"

type node struct {
	id   int
	next *node
}

func (n *node) setNext(next *node) {
	n.next = next
}

func (n *node) print() {
	temp := n

	for ; temp.next != nil; temp = temp.next {
		fmt.Printf("%d -> ", temp.id)
	}

	fmt.Println(temp.id)
}

// 1   2   3   4   5   6
//
// ^
//                 ^   *
//
// head move from 1 to 5(递)
// newHead settled at 6
// head move from 5 to 1(归)

func reverse(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}

	newHead := reverse(head.next)

	head.next.next = head
	head.next = nil

	return newHead
}

func initDemoLinkedList() *node {
	n1 := &node{id: 1}
	n2 := &node{id: 2}
	n3 := &node{id: 3}
	n4 := &node{id: 4}
	n5 := &node{id: 5}
	n6 := &node{id: 6}

	n1.setNext(n2)
	n2.setNext(n3)
	n3.setNext(n4)
	n4.setNext(n5)
	n5.setNext(n6)

	return n1
}

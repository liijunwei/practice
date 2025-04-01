package algo

type Node struct {
	Val  int
	Next *Node
}

func BuildLinkedListFromArray(a []int) *Node {
	if len(a) == 0 {
		return nil
	}

	head := &Node{Val: a[0]}
	curr := head

	for _, val := range a[1:] {
		nextNode := &Node{Val: val}
		curr.Next = nextNode
		curr = nextNode
	}

	return head
}

func BuildCycledLinkedListFromArray(a []int) *Node {
	if len(a) == 0 {
		return nil
	}

	head := &Node{Val: a[0]}
	curr := head

	for _, val := range a[1:] {
		nextNode := &Node{Val: val}
		curr.Next = nextNode
		curr = nextNode
	}

	curr.Next = head

	return head
}

func LinkedListToArray(head *Node) []int {
	if head == nil {
		return []int{}
	}

	arr := make([]int, 0, 100)
	curr := head

	for {
		arr = append(arr, curr.Val)
		curr = curr.Next

		if curr == head {
			break
		}
	}

	return arr
}

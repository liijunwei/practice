package algo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 每隔两个元素删除一个元素 <=> 1，2，3报数，删除报数为3的元素
func TestJosephusProblem1(t *testing.T) {
	head := BuildCycledLinkedListFromArray([]int{1, 2, 3, 4, 5, 6})
	require.True(t, head != nil)

	arr := LinkedListToArray(head)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, arr)

	// fmt.Println("The Last of Us", JosephusProblemSolution1(head, 2))
	// assert.Equal(t, 1, head.Val)

	// TODO this looks wrong
	n := 2000
	longerlist := make([]int, 0, n)
	for i := range n {
		longerlist = append(longerlist, i+1)
	}

	head1 := BuildCycledLinkedListFromArray(longerlist)
	require.True(t, head1 != nil)
	res, val := JosephusProblemSolution1(head1, 2)
	fmt.Println("the order of deletion", res)
	fmt.Println("The Last of Us", val)
}

func JosephusProblemSolution1(list *Node, n int) ([]int, int) {
	curr := &Node{Next: list}
	deleteArr := []int{}

	for curr.Next != nil && curr.Next != curr {
		for range n {
			curr = curr.Next
		}
		// fmt.Println("deleting", curr.Next.Val)
		deleteArr = append(deleteArr, curr.Next.Val)
		curr.Next = curr.Next.Next
	}

	return deleteArr, list.Val
}

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

	// arr := LinkedListToArray(list)
	// assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, arr)

	fmt.Println("The Last of Us", JosephusProblemSolution1(head, 2))
	assert.Equal(t, 1, head.Val)
}

func JosephusProblemSolution1(list *Node, n int) int {
	dummy := &Node{Next: list}
	curr := dummy

	counter := 0
	for curr.Next != nil && curr.Next != curr {
		for range n {
			curr = curr.Next
		}
		fmt.Println("deleting", curr.Next.Val)
		curr.Next = curr.Next.Next

		// detect dead loop
		counter++
		if counter > 500 {
			break
		}
	}

	return list.Val
}

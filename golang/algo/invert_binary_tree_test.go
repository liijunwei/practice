package algo

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/invert-binary-tree/
func TestInvertBinaryTree(t *testing.T) {
	root1 := &TreeNode{Val: 4}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 7}
	root1.Left.Left = &TreeNode{Val: 1}
	root1.Left.Right = &TreeNode{Val: 3}
	root1.Right.Left = &TreeNode{Val: 6}
	root1.Right.Right = &TreeNode{Val: 9}

	fmt.Println("root1", treeToArray(root1))
	invertedRoot1 := invertTree(root1)
	fmt.Println("invertedRoot1", treeToArray(invertedRoot1))

	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}

	fmt.Println("root2", treeToArray(root2))
	invertedRoot2 := invertTree(root2)
	fmt.Println("invertedRoot2", treeToArray(invertedRoot2))

	var root3 *TreeNode
	fmt.Println("root3", treeToArray(root3))
	invertedRoot3 := invertTree(root3)
	fmt.Println("invertedRoot3", treeToArray(invertedRoot3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left

	return root
}

// treeToArray converts a binary tree to an array representation
// using level-order traversal (BFS)
func treeToArray(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

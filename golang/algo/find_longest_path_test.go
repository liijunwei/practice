package algo

import (
	"fmt"
	"testing"
)

//    1
//  2   3
// 4 5 6 7
//  8 9

func Test_findLongestPath(t *testing.T) {
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 5}
	root1.Right.Left = &TreeNode{Val: 6}
	root1.Right.Right = &TreeNode{Val: 7}

	root1.Left.Right.Left = &TreeNode{Val: 8}
	root1.Left.Right.Right = &TreeNode{Val: 9}

	fmt.Println(findLongestPath(root1))
}

func findLongestPath(root *TreeNode) []int {
	var dfs func(*TreeNode)

	path := []int{}
	longestPath := []int{}
	seen := make(map[int]struct{})

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if _, ok := seen[node.Val]; ok {
			return
		}

		seen[node.Val] = struct{}{}
		path = append(path, node.Val)

		// fmt.Println("visiting", node.Val)

		if node.Left == nil && node.Right == nil {
			if len(path) > len(longestPath) {
				longestPath = make([]int, len(path))
				copy(longestPath, path)
				// fmt.Println(longestPath)
			}
		}

		dfs(node.Left)
		dfs(node.Right)

		// 回溯：移除当前节点
		// fmt.Println(path)
		path = path[:len(path)-1]
		delete(seen, node.Val)
	}

	dfs(root)

	return longestPath
}

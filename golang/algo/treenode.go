package algo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
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

func boom(ok bool) {
	if !ok {
		panic("should not happen")
	}
}

func (root *TreeNode) Insert(val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = root.Left.Insert(val)
	} else {
		root.Right = root.Right.Insert(val)
	}

	return root
}

// root -> left -> right
func (root *TreeNode) PreOrder() []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, root.Val)
	result = append(result, root.Left.PreOrder()...)
	result = append(result, root.Right.PreOrder()...)

	return result
}

// left -> root -> right
func (root *TreeNode) InOrder() []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, root.Left.InOrder()...)
	result = append(result, root.Val)
	result = append(result, root.Right.InOrder()...)

	return result
}

func inorderTraversal(root *TreeNode) []int {
	// left->root->right
	var traverse func(root *TreeNode) []int

	traverse = func(root *TreeNode) []int {
		if root == nil {
			return nil
		}

		result := make([]int, 0, 100)
		result = append(result, traverse(root.Left)...)
		result = append(result, root.Val)
		result = append(result, traverse(root.Right)...)

		return result
	}

	return traverse(root)
}

// left -> right -> root
func (root *TreeNode) PostOrder() []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, root.Left.PostOrder()...)
	result = append(result, root.Right.PostOrder()...)
	result = append(result, root.Val)

	return result
}

// LevelOrder traverses the tree in level order (breadth-first)
// and returns an array of values
func (root *TreeNode) LevelOrder() []int {
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

// https://www.hello-algo.com/chapter_tree/array_representation_of_tree/#732
func BuildTreeFromLevelOrder(arr []int, nullValue int) *TreeNode {
	panic("TODO")
}

func (root *TreeNode) Remove(val int) *TreeNode {
	panic("TODO")
}

// Implement the Search function that was marked as TODO
func (root *TreeNode) Search(val int) bool {
	panic("TODO")
}

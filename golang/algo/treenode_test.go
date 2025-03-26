package algo

import (
	"reflect"
	"testing"
)

func TestTreeIteration(t *testing.T) {
	//  2
	// 1  3
	//	    4
	var root1 *TreeNode
	root1 = root1.Insert(2).Insert(3).Insert(1).Insert(4)
	assertEqual(t, root1.PreOrder(), []int{2, 1, 3, 4})
	assertEqual(t, root1.InOrder(), []int{1, 2, 3, 4})
	assertEqual(t, root1.InOrder(), inorderTraversal(root1))
	assertEqual(t, root1.PostOrder(), []int{1, 4, 3, 2})

	//   5
	//  3  7
	// 2  6 8
	var root2 *TreeNode
	root2 = root2.Insert(5).Insert(3).Insert(7).Insert(2).Insert(6).Insert(8)
	assertEqual(t, root2.PreOrder(), []int{5, 3, 2, 7, 6, 8})
	assertEqual(t, root2.InOrder(), []int{2, 3, 5, 6, 7, 8})
	assertEqual(t, root2.InOrder(), inorderTraversal(root2))
	assertEqual(t, root2.PostOrder(), []int{2, 3, 6, 8, 7, 5})
}

func assertEqual(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Fatal("expected", a, "but got", b)
	}
}

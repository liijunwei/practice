package algo

import (
	"testing"
)

func TestBST(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(4)
	bst.Insert(5)
	bst.Insert(6)
	bst.Insert(7)

	// assert.Equal(t, 5, bst.root.Val)
	// assert.Equal(t, 3, bst.root.Left.Val)
	// assert.Equal(t, 7, bst.root.Right.Val)
	// assert.Equal(t, []int{5, 3, 7, 2, 4, 6, 8}, treeToArray(bst.root))

	// bst.Levelorder()
	// spew.Dump(bst)
}

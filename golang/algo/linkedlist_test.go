package algo

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestBuildLinkedListFromArray(t *testing.T) {
	spew.Dump(BuildLinkedListFromArray([]int{1, 2, 3}))
	spew.Dump(BuildLinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7}))
	spew.Dump(BuildCycledLinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7}))
}

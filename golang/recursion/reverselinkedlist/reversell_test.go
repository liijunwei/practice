package reversell

import "testing"

func Test_reverseLinkedList(t *testing.T) {
	head := initDemoLinkedList()
	head.print()

	newHead := reverse(head)
	newHead.print()
}

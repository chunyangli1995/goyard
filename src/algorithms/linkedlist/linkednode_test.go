package linkedlist

import "testing"

func TestLinkedNode(t *testing.T) {
	tail := &LinkedNode{value: 2}
	head := &LinkedNode{2, tail}
	if head.value != 2 || head.next != tail || tail.next != nil {
		t.Error("uncorrect linked node")
	}
}

package linkedlist

import "testing"

func TestLinkedNode(t *testing.T) {
	tail := &LinkedNode{value: 2}
	head := &LinkedNode{2, tail}
	if head.value != 2 || head.next != tail || tail.next != nil {
		t.Error("uncorrect linked node")
	}
}

func TestRemoveNode(t *testing.T) {
	node2 := &LinkedNode{3, nil}
	node1 := &LinkedNode{2, node2}
	head := &LinkedNode{1, node1}
	if head.next.value != 2 || head.next.next.value != 3 {
		t.Error("uncorrect linked list")
	}
	RemoveNode(head, 2)
	if head.next.value != 3 || head.next.next != nil {
		t.Error("uncorrect linked list")
	}
}

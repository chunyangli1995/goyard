package tree

import "testing"

func TestConstruct(t *testing.T) {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	rootNode, e := Construct(preOrder, inOrder)
	if e != nil {
		t.Error(e.Error())
	}
	if rootNode.value != 1 || rootNode.leftChild.value != 2 || rootNode.leftChild.leftChild.value != 4 ||
		rootNode.leftChild.leftChild.rightChild.value != 7 || rootNode.rightChild.value != 3 ||
		rootNode.rightChild.leftChild.value != 5 || rootNode.rightChild.rightChild.value != 6 ||
		rootNode.rightChild.rightChild.leftChild.value != 8 {
		t.Error("uncorrect binary tree construct")
	}
}

func TestVerifySequenceOfBst(t *testing.T) {
	data := []int{5, 7, 6, 9, 11, 10, 8}
	ret := VerifySequenceOfBst(data)
	if ret == false {
		t.Error("uncorrect ret.")
	}
}

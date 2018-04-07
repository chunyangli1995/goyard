package tree

import (
	"container/list"
	goyardError "goyard/src/err"
)

type BinaryTreeNode struct {
	value      int
	leftChild  *BinaryTreeNode
	rightChild *BinaryTreeNode
}

func constructCore(preOrder []int, inOrder []int) (*BinaryTreeNode, error) {
	rootValue := preOrder[0]
	rootNode := new(BinaryTreeNode)
	rootNode.value = rootValue
	if len(preOrder) == 1 {
		return rootNode, nil
	}
	rootInorderIndex := 0
	for rootInorderIndex < len(inOrder) && inOrder[rootInorderIndex] != rootValue {
		rootInorderIndex++
	}
	if rootInorderIndex == len(inOrder)-1 && inOrder[rootInorderIndex] != rootValue {
		return nil, goyardError.NewInvalidInputError()
	}
	if rootInorderIndex >= 1 {
		rootNode.leftChild, _ = constructCore(preOrder[1:1+rootInorderIndex], inOrder[0:rootInorderIndex])
	}
	if rootInorderIndex < len(inOrder)-1 {
		rootNode.rightChild, _ = constructCore(preOrder[rootInorderIndex+1:], inOrder[rootInorderIndex+1:])
	}
	return rootNode, nil
}

/**
根据前序遍历和中序遍历还原二叉树
*/
func Construct(preOrder []int, inOrder []int) (*BinaryTreeNode, error) {
	if preOrder == nil || inOrder == nil || len(inOrder) != len(preOrder) || len(inOrder) == 0 {
		return nil, goyardError.NewInvalidInputError()
	}
	return constructCore(preOrder, inOrder)
}

func DoesTree1HasTree2(tree1 *BinaryTreeNode, tree2 *BinaryTreeNode) bool {
	/**
	判断 tree1 是否包含 tree2
	*/
	if tree1 == nil {
		return false
	}
	if tree2 == nil {
		return true
	}
	if tree1.value != tree2.value {
		return false
	}
	return DoesTree1HasTree2(tree1.leftChild, tree2.leftChild) && DoesTree1HasTree2(tree1.rightChild, tree2.rightChild)
}

func HasSubTree(tree1 *BinaryTreeNode, tree2 *BinaryTreeNode) bool {
	result := false
	if tree1 != nil && tree2 != nil {
		if tree1.value == tree2.value {
			result = DoesTree1HasTree2(tree1, tree2)
		}
		if result == false {
			result = DoesTree1HasTree2(tree1.leftChild, tree2)
		}
		if result == false {
			result = DoesTree1HasTree2(tree1.rightChild, tree2)
		}
	}
	return result
}

func VerifySequenceOfBst(sequence []int) bool {
	if sequence == nil || len(sequence) == 0 {
		return false
	}
	root := sequence[len(sequence)-1]
	i := 0
	for i < len(sequence)-1 {
		if sequence[i] > root {
			break
		}
		i++
	}
	j := i
	for j < len(sequence)-1 {
		if sequence[j] < root {
			return false
		}
		j++
	}
	left := true
	if i > 0 {
		left = VerifySequenceOfBst(sequence[:i])
	}
	right := true
	if i < len(sequence)-1 {
		right = VerifySequenceOfBst(sequence[i : len(sequence)-1])
	}
	return left && right
}

func Convert(root *BinaryTreeNode) *BinaryTreeNode {
	var lastNodeInList *BinaryTreeNode
	convertNode(root, lastNodeInList)
	for lastNodeInList != nil && lastNodeInList.leftChild != nil {
		lastNodeInList = lastNodeInList.leftChild
	}
	return lastNodeInList
}

func convertNode(root *BinaryTreeNode, lastNodeInList *BinaryTreeNode) {
	if root == nil {
		return
	}
	current := root
	if current.leftChild != nil {
		convertNode(root.leftChild, lastNodeInList)
	}
	current.leftChild = lastNodeInList
	if lastNodeInList != nil {
		lastNodeInList.rightChild = current
	}
	lastNodeInList = current
	if current.rightChild != nil {
		convertNode(root.rightChild, lastNodeInList)
	}
}

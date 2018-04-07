package linkedlist

type LinkedNode struct {
	value int
	next  *LinkedNode
}

type ComplexLinkedNode struct {
	value   int
	next    *ComplexLinkedNode
	sibling *ComplexLinkedNode
}

func RemoveNode(head *LinkedNode, value int) {
	if head == nil {
		return
	}
	if head.value == value {
		head = head.next
	} else {
		pNode := head
		for pNode.next != nil && pNode.next.value != value {
			pNode = pNode.next
		}
		if pNode.next != nil && pNode.next.value == value {
			pNode.next = pNode.next.next
		}
	}
}

func DeleteNode(head *LinkedNode, deleted *LinkedNode) {
	if head == nil || deleted == nil {
		return
	}
	if deleted.next != nil {
		deleted.value = deleted.next.value
		deleted.next = deleted.next.next
	} else if head == deleted {
		deleted = nil
		head = nil
	} else {
		pNode := head
		for pNode.next != deleted {
			pNode = pNode.next
		}
		pNode.next = nil
		deleted = nil
	}
}

func FindKthNodeToTail(head *LinkedNode, k int) *LinkedNode {
	if head == nil || k <= 0 {
		return nil
	}
	pHead := head
	for i := 0; i < k; i++ {
		if pHead.next != nil {
			pHead = pHead.next
		} else {
			return nil
		}
	}
	pBehind := head
	for pHead.next != nil {
		pHead = pHead.next
		pBehind = pBehind.next
	}
	return pBehind
}

func ReverseList(head *LinkedNode) *LinkedNode {
	var reversedHead *LinkedNode
	pNode := head
	var pPrev *LinkedNode
	for pNode != nil {
		pNext := pNode.next
		if pNext == nil {
			reversedHead = pNode
		}
		pNode.next = pPrev
		pPrev = pNode
		pNode = pNext
	}
	return reversedHead
}

func CloneNodes(head *ComplexLinkedNode) {
	node := head
	for node != nil {
		cloneNode := &ComplexLinkedNode{}
		cloneNode.value = node.value
		cloneNode.next = node.next
		node.next = cloneNode
		node = cloneNode.next
	}
}

func ConnectSiblingNodes(head *ComplexLinkedNode) {
	node := head
	for node != nil {
		cloneNode := node.next
		if node.sibling != nil {
			cloneNode.sibling = node.sibling.next
		}
		node = cloneNode.next
	}
}

func ReConnectNodes(head *ComplexLinkedNode) *ComplexLinkedNode {
	node := head
	var cloneHead *ComplexLinkedNode
	var cloneNode *ComplexLinkedNode
	if node != nil {
		cloneHead = node.next
		cloneNode = node.next
		node.next = cloneNode.next
		node = node.next
	}
	for node != nil {
		cloneNode.next = node.next
		cloneNode = cloneNode.next
		node.next = cloneNode.next
		node = node.next
	}
	return cloneHead
}

func CloneComplexNode(head *ComplexLinkedNode) *ComplexLinkedNode {
	CloneNodes(head)
	ConnectSiblingNodes(head)
	return ReConnectNodes(head)
}

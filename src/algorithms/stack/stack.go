package stack

import "container/list"

func IsPopOrder(pPush []int, pPop []int, nLength int) bool {
	bPossible := false
	if pPush != nil && pPop != nil && nLength > 0 {
		pushIndex := 0
		popIndex := 0
		stackData := list.New()
		for popIndex < nLength {
			for stackData.Len() == 0 || stackData.Front().Value != pPop[popIndex] {
				if popIndex >= nLength {
					break
				}
				stackData.PushFront(pPush[pushIndex])
				pushIndex++
			}
			if stackData.Front().Value != pPop[popIndex] {
				break
			}
			stackData.Remove(stackData.Front())
			popIndex++
		}
		if stackData.Len() == 0 && popIndex == nLength {
			bPossible = true
		}
	}
	return bPossible
}

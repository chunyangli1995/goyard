package stack

import "testing"

func TestIsPopOrder(t *testing.T) {
	pushOrder := []int{1, 2, 3, 4, 5}
	popOrder := []int{4, 5, 3, 2, 1}
	bPossible := IsPopOrder(pushOrder, popOrder, len(pushOrder))
	if bPossible == false {
		t.Error("uncorrect ret.")
	}
}

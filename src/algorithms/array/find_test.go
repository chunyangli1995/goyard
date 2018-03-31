package array

import "testing"

func TestFind(t *testing.T) {
	a := [][]int{
		{1, 4, 5, 7},
		{2, 6, 8, 9},
		{3, 10, 11, 12},
	}
	found := FindElemInMatrix(a, 3, 4, 10)
	if found == false {
		t.Error("elem not found in matrix")
	}
	found = FindElemInMatrix(a, 3, 4, 13)
	if found == true {
		t.Error("unexist elem found in matrix")
	}
}

package search

import "testing"

func TestSearch(t *testing.T) {
	data := []int{3, 4, 5, 1, 2}
	minVal := min(data)
	if minVal != 1 {
		t.Error("uncorrect min val in list")
	}
}

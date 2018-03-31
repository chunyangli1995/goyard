package array

/**
数组从左到右变大，从上到下变大
*/
func FindElemInMatrix(a [][]int, rows, cols, value int) bool {
	found := false
	if a != nil && rows > 0 && cols > 0 {
		row := 0
		col := cols - 1
		for row < rows && col >= 0 {
			if a[row][col] == value {
				found = true
				break
			} else if a[row][col] > value {
				col -= 1
			} else {
				row += 1
			}
		}
	}
	return found
}

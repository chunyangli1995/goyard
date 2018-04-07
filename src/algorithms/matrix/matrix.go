package matrix

func PrintMatrixClockwisely(numbers [][]int, columns int, rows int) {
	if numbers == nil || columns <= 0 || rows <= 0 {
		return
	}
	start := 0
	for columns > start*2 && rows > start*2 {
		printMatrixInCircle(numbers, columns, rows, start)
		start++
	}
}

func printMatrixInCircle(numbers [][]int, columns int, rows int, start int) {
	endX := columns - 1 - start
	endY := rows - 1 - start
	// 左到右
	for i := start; i <= endX; i++ {
		print(numbers[start][i], " ")
	}
	// 上到下
	if start < endY {
		for i := start + 1; i <= endY; i++ {
			print(numbers[i][endY], " ")
		}
	}
	// 右到左
	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			print(numbers[endY][i], " ")
		}
	}
	// 下到上
	if start < endX && start < endY-1 {
		for i := endY - 1; i >= start+1; i-- {
			print(numbers[i][start], " ")
		}
	}
}

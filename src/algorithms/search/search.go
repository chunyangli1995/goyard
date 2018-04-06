package search

/**
一个从小到大排序的整数数组，把前面若干位移动到后面，求数组的最小值
思路：二分
类似变种：升序降序数组找最大值，降序升序数组找最小值，本质都是通过二分
*/
func min(numbers []int) int {
	if numbers == nil || len(numbers) == 0 {
		panic("invalid input")
	}
	index1 := 0
	index2 := len(numbers) - 1
	mid := index1
	for numbers[index1] >= numbers[index2] {
		if index2-index1 == 1 {
			return numbers[index2]
		}
		mid = (index1 + index2) / 2
		if numbers[index1] == numbers[index2] && numbers[index1] == numbers[mid] {
			return minInOrder(numbers, index1, index2)
		}
		if numbers[mid] >= numbers[index1] {
			index1 = mid
		}
		if numbers[mid] <= numbers[index2] {
			index2 = mid
		}
	}
	return numbers[index1]

}

func minInOrder(numbers []int, index1 int, index2 int) int {
	minVal := numbers[index1]
	for i := index1 + 1; i <= index2; i++ {
		if numbers[i] < minVal {
			minVal = numbers[i]
		}
	}
	return minVal
}

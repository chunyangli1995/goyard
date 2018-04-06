package recursive

func PrintToMaxOfNdigitsRecursively(numbers [10]int, length int, index int) {
	if index == length-1 {
		printNumber(numbers, length)
		return
	}
	for i := 0; i < 10; i++ {
		numbers[index+1] = i
		PrintToMaxOfNdigitsRecursively(numbers, length, index+1)
	}
}

func printNumber(numbes [10]int, length int) {
	/**
	打印数组，开头的0需要忽略
	*/
	isBegin := false
	for i := 0; i < length; i++ {
		if isBegin == false && numbes[i] != 0 {
			isBegin = true
		}
		if isBegin == true {
			print(numbes[i])
		}
	}
	println()
}

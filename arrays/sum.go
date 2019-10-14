package main

// Sum calc sum
func Sum(numbers []int) int {
	sum := 0
	// index and value
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll ... 表示多个同类型参数
func SumAll(numbersToSum ...[]int) []int {
	// 可以使用 append 函数为切片追加一个新值
	var sums []int
	// 可以使用 = 对切片的元素赋值
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails 将尾部元素相加
func SumAllTails(numbersToSum ...[]int) []int {
	// 可以使用 append 函数为切片追加一个新值
	var sums []int
	// 可以使用 = 对切片的元素赋值
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {

			// 提取出尾部元素
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

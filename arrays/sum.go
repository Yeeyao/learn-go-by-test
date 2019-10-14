package main

// // Sum calc sum
// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

// Sum calc sum using range
// 注意，重构之后，需要维护测试代码

/*
质疑测试的价值是非常重要的。测试不是越多越好，而是要尽可能的使你的代码更加健壮
太多的测试会增加维护成本，因为维护每个测试都是需要成本的
*/

/*
	len 获取数组和切片的长度
	append 在切片尾部添加元素
	go test -cover 测试覆盖率
	reflect.DeepEqual 这个判断函数不是类型安全的
	切片中的 [n:m] 是前闭后开的
*/

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

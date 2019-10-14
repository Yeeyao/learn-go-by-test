package main

import "testing"

func TestSum(t *testing.T) {

	// 数组容量是我们在声明它们时指定的固定值
	// 初始化时，[] 内是数字则是显示指定容量，是 ... 则不指定容量，让编译器自行根据初始列表计算容量
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

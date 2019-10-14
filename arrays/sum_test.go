package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		// 数组容量是我们在声明它们时指定的固定值
		// 初始化时，[] 内是数字则是显示指定容量，是 ... 则不指定容量，让编译器自行根据初始列表计算容量
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		// no
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// want := "bob"

	// 这个判断函数不是类型安全的
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// TestSumAllTails 把切片的尾部元素相加
func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("make the sums of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

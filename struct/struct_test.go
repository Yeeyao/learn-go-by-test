package main

import "testing"

/*
方法

当我们调用t.Errorf时，我们调用了 t(testing.T) 这个实例的方法 Errorf
方法和函数很相似，但是方法是通过一个特定类型的示例调用的，函数可以随时被调用

声明方法的语法类似函数，唯一不同的是方法的接收者的语法
func(receiverName ReceiverType) MethodName(args)
方法被这种类型的变量调用时，数据的引用通过变量 receiverName 获得
把类型的第一个字母作为接收者变量是 Go 语言的一个惯例

接口

Go 这种静态类型语言中的一种非常强有力的概念，
因为接口可以让函数接受不同类型的参数并能够创造类型安全且高解耦的代码

*/

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area(rectangle)
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area(circle)
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

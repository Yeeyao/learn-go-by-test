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

接口指定了一个方法集合
一个接口类型的变量可以存储任何含有接口方法集合的超集的类型的一个值

在 Go 语言中，interface resolution 是隐式的，如果传入的类型匹配接口的需要，则编译正确

这里的例子

Rectangle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
Circle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
引入新的几何形状只需要实现 Area 方法并把新的类型加到测试用例中

列表驱动测试，在测试中真正需要使用它，比如测试一个接口的不同表现，传入函数数据有很多不同的测试需求

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

	// 匿名结构体以及声明了一个结构体切片
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}

	}

}

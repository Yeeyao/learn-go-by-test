package main

import "testing"

/*

   当我们不知道要编写的函数参数在编译时是什么类型的时候
   Go 允许我们使用 interface{} 来解决，可以将它视为任意类型
   函数的参数定义为 interface{} 类型时，必须检查所有传入的参数
   并断定参数类型以及如何处理它们，这是通过反射实现的。
   除非真的需要，一般不使用反射。

   如果想实现函数的多态性，考虑是否可以围绕接口。

*/

func TestWalk(t *testing.T) {

	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got '%s' want '%s'", got, expected)
	}
}

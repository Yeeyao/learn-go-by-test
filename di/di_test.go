package main

import (
	"bytes"
	"testing"
)

/*
测试打印数据的例子
	我们的函数不需要关心在哪里打印以及如何打印，所以我们应该接收一个接口，而非一个具体的类型

bytes 包中的 buffer 类型实现了 writer 接口，因此我们可以用它作为我们的 Writer，
调用 Greet 后，我们可以用它来检查写入了什么
io.Writer 是一个很好的通用接口，用于将数据存放到某个地方

fmt.Fprintf 会接收一个 Writer 参数，用于把字符串传递过去

*/

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

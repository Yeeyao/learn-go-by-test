package main

import (
	"bytes"
	"testing"
)

/*
测试打印数据的例子
	我们的函数不需要关心在哪里打印以及如何打印，所以我们应该接收一个接口，而非一个具体的类型

bytes 包中的 Buffer 类型实现了 writer 接口，因此我们可以用它作为我们的 Writer，
调用 Greet 后，我们可以用它来检查写入了什么
bytes.Buffer 实现了 io.Writer 接口
io.Writer 是一个很好的通用接口，用于将数据存放到某个地方

记住 fmt.Fprintf 和 fmt.Printf 一样，
只不过 fmt.Fprintf 会接收一个 Writer 参数，用于把字符串传递过去，而 fmt.Printf 默认是标准输出。

注入依赖 dependency injection，我们可以控制数据向哪里写入
	同时允许我们测试代码 DI 提倡注入了一个数据库依赖（通过接口），就可以在测试中控制模拟数据
	关注点分离 解耦了数据到达的地方和如何产生数据。当一个方法 / 函数负责太多功能了
	在不同环境下重用代码

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

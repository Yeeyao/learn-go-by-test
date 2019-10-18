package main

import (
	"bytes"
	"testing"
)

/*
我们有一个 Sleep ing 的注入，需要抽离出来我们才能在测试中控制它
如果我们能够 mock time.Sleep，我们可以用依赖注入的方式去代替真正的
times.Sleep，然后我们可以使用断言监视调用

监视器（spies）是一种 mock，它可以记录依赖关系是怎样被使用的，传入的参数
，多少次等等。（我们的例子中记录 Sleep() 被调用了多少次）
*/
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}

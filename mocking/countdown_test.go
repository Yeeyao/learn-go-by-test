package main

import (
	"bytes"
	"reflect"
	"testing"
)

/*
我们有一个 Sleep ing 的注入，需要抽离出来我们才能在测试中控制它
如果我们能够 mock time.Sleep，我们可以用依赖注入的方式去代替真正的
times.Sleep，然后我们可以使用断言监视调用

监视器（spies）是一种 mock，它可以记录依赖关系是怎样被使用的，传入的参数
，多少次等等。（我们的例子中记录 Sleep() 被调用了多少次）

mocking evil?
如果你的模拟代码变得很复杂，或者需要模拟很多东西来测试一些东西，那么你就应该考虑
	你正在进行的测试需要做太多的事情
	它的依赖关系太细致了
	你的测试过于关注实现细节

测试驱动开发，通常情况下，糟糕的测试代码是糟糕的设计的结果，而设计良好的代码和容易测试
需要注意测试到什么级别 测试我想要的行为还是实现细节 重构的代码，需要对测试进行很多修改
小心使用监视器 这将使得实现细节和代码之前耦合更加紧密

测试驱动开发的方法
	把问题分解为简单的模块，试着让你的工作软件尽快得到测试的支持，以避免掉进（rabbit holes，指未知领域）和采取最终测试的方法
	一旦有一些正在工作的软件，小步迭代应该是很容易的，直到你实现你所需要的软件

Mocking
	没有对代码中重要的取悦进行 mock 将会导致难以测试。
	没有 mock，可能需要设置一些其他东西来测试简单的业务规则，导致缓慢的反馈循环
	服务的不可靠性将得到一个脆弱的测试
	避免过度测试，始终注意测试的价值以及它们在将来的重构中会产生什么样的影响

*/

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

package main

import (
	"testing"
)

/*

	这里，主程序文件中需要有package
	Go 中，如果一个符号（变量，类型，函数等）是以小写符号开头，那么它在定义它的包之外是私有的

	Go 允许从现有的类型创建新的类型
	type MyName OriginalType

	类型别名有一个有趣的特性，可以对它们声明方法，可以用来在现有类型之上添加一些领域内特定的功能

	nil 是其他编程语言的 null。错误可以是 nil，因为返回类型是 error，这是一个接口。
	一个函数的参数或者返回值的类型是一个接口，他们就可以是 nil。

	Go 中，错误是值，因此我们可以将其重构为一个变量，并为其提供一个单一的事实来源
	即单一的事实来源方便定位错误
	这里将助手函数从主测试函数中移除，他人打开文件时就可以读取断言
	测试的另一个有用的特性是，它帮主我们理解代码的真实用途，从而使我们的代码更具有交互性

	未经检查的错误
	可以在终端中运行 errcheck 来安装该 linter
	之后可以在代码目录下运行 errcheck .

	指针

	如果函数需要改变状态，就需要用指针指向你想要更改的值，同时，也可以传递一个引用
	当函数返回一个指针时，需要检查它是否为 nil

	错误

	在调用函数或者方法时表示失败

*/

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		// 初始化 20 个
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		// 使用 var
		assertError(t, err, InsufficientFundsError)
	})

}

// 快速测试的助手方法
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

// 快速测试的助手方法
// t.Fatal 如果被调用，将停止测试
func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

// 快速测试的助手方法
func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but want one")
	}

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

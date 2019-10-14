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

*/

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error) {
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		// 初始化 20 个
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		got := wallet.Balance()
		assertBalance(t, wallet, got)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})
}

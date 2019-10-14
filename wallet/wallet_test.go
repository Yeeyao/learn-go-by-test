package main

import (
	"testing"
)

/*

	这里，主程序文件中需要有package
	Go 中，如果一个符号（变量，类型，函数等）是以小写符号开头，那么它在定义它的包之外是私有的

	Go 允许从现有的类型创建新的类型

*/

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

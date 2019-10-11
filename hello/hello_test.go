package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// 这里告诉测试套件这个方法是辅助函数，测试失败时报告的行号是函数调用内而不是辅助函数中
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "french")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

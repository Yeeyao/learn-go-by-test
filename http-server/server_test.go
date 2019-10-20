package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	就 TDD 的安全性而言，编写最少量的代码规则十分重要。
	测试未通过前你写得越多，也就引入越来越多测试不能覆盖的问题。

	模拟测试
	可以创建一个简单的存根来测试代码而无需实现任何真实的存储机制

	可以先写一个硬编码的值让测试通过来启动我们的工作，一旦有了可以通过
	测试的用例，我们就可以接着写更多测试来帮我们删除之前的硬编码代码。
*/

func TestGetPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}

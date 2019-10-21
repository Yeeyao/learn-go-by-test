package main

import (
	"fmt"
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
	store := StubPlayerStore{
		map[string]int {
			"Pepper": 20,
			"Floyd": 10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")

	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}

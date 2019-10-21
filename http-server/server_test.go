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

// 为何这部分必须要定义在这里才能通过测试，定义在 server.go 则不行？
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")

	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}

	// 旧代码使用的是 MethodGet ，这里使用 MethodPost
	//t.Run("it returns accepted on POST", func(t *testing.T) {
	//	request := NewPostWinRequest("Pepper")
	//	response := httptest.NewRecorder()
	//
	//	server.ServeHTTP(response, request)
	//
	//	assertStatus(t, response.Code, http.StatusAccepted)
	//
	//	if len(store.winCalls) != 1 {
	//		t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	//	}
	//})

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], player)
		}
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("dit not get correct status, got %d want %d", got, want)
	}
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

func NewPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

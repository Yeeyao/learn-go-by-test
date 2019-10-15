package main

import "testing"

func TestSearch(t *testing.T) {
	// [] 内是 key 类型，后面跟着的是 value 类型
	// 这里 key 类型必须是可比较类型
	// 通过 map[key] 获取 map 的值
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

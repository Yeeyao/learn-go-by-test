package concurrency

import (
	"reflect"
	"testing"
)

/*
Go 中不会阻塞的操作将在称为 goroutine 的单独的进程中进行。

匿名函数，可以在声明的同时执行，匿名函数末尾 () 来实现。维护对其所定义的词汇作用域的访问权

channels
一个 Go 数据结构，可以同时接收和发送值。这些操作以及细节允许不同进程间的通信

*/

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{

		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}

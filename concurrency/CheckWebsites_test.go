package concurrency

import (
	"reflect"
	"testing"
)

/*
Go 中不会阻塞的操作将在称为 goroutine 的单独的进程中进行。
*/

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckwebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http:blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{

		"http://google.com":        true,
		"http:blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":  false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}

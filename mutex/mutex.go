package main

import (
	"fmt"
	"sync"
)

var x = 0

// using buffered channel with capacity 1
// channel and mutex 都可以使用，取决于问题对哪个适用
// 一般在 Goroutines 间需要通信时使用 channels
func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<- ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

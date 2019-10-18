package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {

	/*
		select 允许你同时在多个 channel 等待，第一个发送值的 channel 将导致其 case 的代码被执行
		time.After 会在你定义的时间过后发送一个新号给 channel 并返回一个 chan 类型，设置超时？
		设置超时来避免阻塞

		httptest
		一种方便创建测试服务器的方法，可以进行可靠和可控的测试
		使用 net/http 相同的接口作为真实的服务器回合真实环境保持一致
	*/
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

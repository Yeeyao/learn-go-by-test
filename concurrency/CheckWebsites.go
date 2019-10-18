package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// chan result 是 channel 类型的 -- result 的 channel
	resultChannel := make(chan result)

	/*
		这里循环每次迭代都会启动一个新的 goroutine，与当前进程（ WebsiteChecker 函数）同时发生，
		每个循环都会将结果添加到 results map 中
		现在，我们迭代 urls 时，不是直接写入 map，而是使用 send statement 将每个
		调用 wc 的 result 结构体发送到 resultChannel。这里使用 <- 操作符，
		channel 放在左边，值放在右边
	*/

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// 这个 for 循环为每个 url 迭代一次，使用 receive expression，它将从通道
	// 接收到的值分配给变量，使用 <- 操作符
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}

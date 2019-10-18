package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// 这里循环每次迭代都会启动一个新的 goroutine，与当前进程（ WebsiteChecker 函数）同时发生，
	// 每个循环都会将结果添加到 results map 中
	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}

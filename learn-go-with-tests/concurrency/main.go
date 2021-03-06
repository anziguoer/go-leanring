package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		// 异步执行
		go func() {
			results[url] = wc(url)
		}()
	}
	time.Sleep(2 * time.Microsecond)
	return results
}

package concurrency

import "time"

// WebsiteChecker is a custom type
type WebsiteChecker func(string) bool

// CheckWebsites recieves WebsiteChecker and array of url, and returns map of status
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}

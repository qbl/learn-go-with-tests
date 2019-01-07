package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func MockWebsiteChecker(url string) bool {
	if url == "invalid://url.format" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"invalid://url.format",
	}

	expected := map[string]bool{
		"http://google.com":    true,
		"invalid://url.format": false,
	}

	actual := CheckWebsites(MockWebsiteChecker, websites)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("actual: %v; expectd: %v", actual, expected)
	}
}

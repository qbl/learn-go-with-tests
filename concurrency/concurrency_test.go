package concurrency

import (
	"reflect"
	"testing"
)

func MockWebsiteChecker(url string) bool {
	if url == "invalid://url.format" {
		return false
	}
	return true
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

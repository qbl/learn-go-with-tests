package racer

import (
	"net/http"
	"time"
)

// Racer accepts two URL string, and returns the quickest URL
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)

	return time.Since(start)
}

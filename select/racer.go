package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {

	aDuartion := measureResponseTime(a)
	bDuartion := measureResponseTime(b)

	if aDuartion < bDuartion {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration{
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

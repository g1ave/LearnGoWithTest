package webRacer

import (
	"net/http"
	"time"
)

func Racer(u1, u2 string) string {
	u1Duration, u2Duration := measureResponseTime(u1), measureResponseTime(u2)
	if u1Duration > u2Duration {
		return u2
	}
	return u1
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

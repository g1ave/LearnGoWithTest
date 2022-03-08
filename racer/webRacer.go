package webRacer

import (
	"net/http"
	"time"
)

func Racer(u1, u2 string) string {
	startA := time.Now()
	http.Get(u1)
	u1Duration := time.Since(startA)

	startB := time.Now()
	http.Get(u2)
	u2Duration := time.Since(startB)

	if u1Duration > u2Duration {
		return u2
	}
	return u1
}

package webRacer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayServer(20 * time.Millisecond)
	fastSever := makeDelayServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastSever.Close()

	slowURL := slowServer.URL
	fastURL := fastSever.URL

	got := Racer(slowURL, fastURL)
	want := fastURL

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

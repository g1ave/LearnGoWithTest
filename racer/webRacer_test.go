package webRacer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastSever := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	slowURL := slowServer.URL
	fastURL := fastSever.URL

	got := Racer(slowURL, fastURL)
	want := fastURL

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
	slowServer.Close()
	fastSever.Close()
}

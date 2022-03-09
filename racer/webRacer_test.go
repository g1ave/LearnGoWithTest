package webRacer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("fast server win", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, _ := Racer(slowURL, fastURL)
		want := fastURL

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

	t.Run("returns an error if a server doesn't response within 10s", func(t *testing.T) {
		slowServer := makeDelayServer(12 * time.Second)
		fastServer := makeDelayServer(11 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, error := Racer(slowURL, fastURL)
		if error == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

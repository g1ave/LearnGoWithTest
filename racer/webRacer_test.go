package webRacer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the URL of the faster one", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, err := Racer(slowURL, fastURL, 10*time.Millisecond)
		want := fastURL

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}

		if err != nil {
			t.Errorf("didn't expect an error but got one")
		}
	})

	t.Run("returns an error if a server doesn't response within 10s", func(t *testing.T) {
		timeoutServer := makeDelayServer(30 * time.Millisecond)

		defer timeoutServer.Close()

		_, err := Racer(timeoutServer.URL, timeoutServer.URL, 20*time.Millisecond)
		if err == nil {
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

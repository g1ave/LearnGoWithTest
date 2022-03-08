package webRacer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://www.facebook.com"
	fastURL := "https://www.baidu.com"

	got := Racer(slowURL, fastURL)
	want := fastURL

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}

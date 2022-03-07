package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteCheck(url string) bool {
	if url[:4] != "http" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"http://google.com",
		"waat://test.com",
		"http://baidu.com",
		"http://sougou.com",
	}

	got := CheckWebsites(mockWebsiteCheck, urls)
	want := map[string]bool{
		"http://google.com": true,
		"waat://test.com":   false,
		"http://baidu.com":  true,
		"http://sougou.com": true,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}

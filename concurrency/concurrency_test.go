package concurrency

import (
	"reflect"
	"testing"
	"time"
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

// Benchmark test
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Microsecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

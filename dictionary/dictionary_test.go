package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is a test"}

	got := Search(dictionary, "test")
	want := "this is a test"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}

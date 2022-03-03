package dictionary

import "testing"

func TestSearch(t *testing.T) {

	t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}

		got, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}

		assertStrings(t, got, "")
	})
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}

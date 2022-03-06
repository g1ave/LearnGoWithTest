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

		assertError(t, err, NotFoundError)

		assertStrings(t, got, "")
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dicionary := Dictionary{}

		word := "test"
		definition := "this is a test"

		addError := dicionary.Add(word, definition)

		got, err := dicionary.Search("test")

		assertNoError(t, addError)
		assertNoError(t, err)
		assertStrings(t, got, definition)

	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		err := dictionary.Add("test", "this is a test")

		assertError(t, err, WordExistedError)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		err := dictionary.Update("test", "new definition")

		got, _ := dictionary.Search("test")
		want := "new definition"

		assertStrings(t, got, want)
		assertNoError(t, err)
	})

	t.Run("update unknown word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("test", "new definition")

		assertError(t, err, WordDoesNotExistError)
	})
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("didn't want an error but got one")
	}
}

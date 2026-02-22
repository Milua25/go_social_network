package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("search for a key", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is just a test"}

		got, err := dictionary.Search("test")
		assertNoError(t, err)
		want := "This is just a test"

		assertStrings(t, got, want)
	})

	t.Run("search for a non-existent key", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is just a test"}
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})

	t.Run("Add a new key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "This is just a test")
		assertNoError(t, err)
		assertDefinition(t, dictionary, "test", "This is just a test")
	})

	t.Run("Add a new key with existing key", func(t *testing.T) {
		word := "test"
		definition := "This is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrAlreadyExists)
		assertDefinition(t, dictionary, word, "This is just a test")
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("no error expected but got:", err)
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got error %q want %q", got, want)
	}
}

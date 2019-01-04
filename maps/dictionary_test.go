package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("knwon word", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertString(t, actual, expected)
	})

	t.Run("unknwon word", func(t *testing.T) {
		_, actual := dictionary.Search("unknown")

		assertError(t, actual, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	word := "test"
	definition := "this is just a test"
	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word, definition)
}

func assertString(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual: %s; expected: %s", actual, expected)
	}
}

func assertError(t *testing.T, actual, expected error) {
	t.Helper()

	if actual != expected {
		t.Errorf("actual error: %s; expected: %s", actual, expected)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	actual, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertString(t, actual, definition)
}

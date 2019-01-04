package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("knwon word", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertString(t, actual, expected)
	})

	t.Run("unknwon word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertString(t, err.Error(), expected)
	})
}

func assertString(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual: %s; expected: %s", actual, expected)
	}
}

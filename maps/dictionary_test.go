package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	actual := Search(dictionary, "test")
	expected := "this is just a test"
	assertString(t, actual, expected)
}

func assertString(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual: %s; expected: %s", actual, expected)
	}
}

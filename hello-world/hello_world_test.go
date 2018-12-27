package main

import "testing"

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld()
	expected := "Hello World!"

	if actual != expected {
		t.Errorf("actual: '%s'; expected: `%s`", actual, expected)
	}
}

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, actual, expected string) {
		t.Helper()
		if actual != expected {
			t.Errorf("actual: '%s'; expected: `%s`", actual, expected)
		}
	}

	t.Run("Hello with non empty parameter", func(t *testing.T) {
		actual := Hello("name", "")
		expected := "Hello name!"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Hello with empty parameter", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello!"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Hello in Spanish", func(t *testing.T) {
		actual := Hello("name", "spanish")
		expected := "Hola name!"
		assertCorrectMessage(t, actual, expected)
	})
}

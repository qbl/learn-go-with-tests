package main

import "testing"

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld()
	expected := "Hello World!"

	if actual != expected {
		t.Errorf("actual: '%s'; expected: `%s`", actual, expected)
	}
}

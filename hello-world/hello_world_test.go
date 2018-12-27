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
	actual := Hello("name")
	expected := "Hello name!"

	if actual != expected {
		t.Errorf("actual: '%s'; expected: `%s`", actual, expected)
	}
}

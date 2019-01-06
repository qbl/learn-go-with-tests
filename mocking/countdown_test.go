package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}

	Countdown(&buffer)

	actual := buffer.String()
	expected := "3"

	if actual != expected {
		t.Errorf("actual: %s; expected: %s", actual, expected)
	}
}

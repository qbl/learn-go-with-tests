package iteration

import "testing"

func TestRepeat(t *testing.T) {
	actual := Repeat("a")
	expected := "aaaaa"

	if actual != expected {
		t.Errorf("expected: %s; actual: %s", expected, actual)
	}
}

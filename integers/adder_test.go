package integers

import "testing"

func TestAdder(t *testing.T) {
	actual := Add(2, 2)
	expected := 4

	if actual != expected {
		t.Errorf("expcted: %d; actual: %d", expected, actual)
	}
}

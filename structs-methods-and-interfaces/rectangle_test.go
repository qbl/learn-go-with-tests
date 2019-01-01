package rectangle

import "testing"

func TestPerimeter(t *testing.T) {
	actual := Perimeter(10.0, 10.0)
	expected := 40.0

	if actual != expected {
		t.Errorf("expcted: %f; actual: %f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	actual := Area(12.0, 6.0)
	expected := 72.0

	if actual != expected {
		t.Errorf("expcted: %f; actual: %f", expected, actual)
	}
}

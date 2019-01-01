package rectangle

import "testing"

var rectangle = Rectangle{10.0, 10.0}

func TestPerimeter(t *testing.T) {
	actual := Perimeter(rectangle)
	expected := 40.0

	if actual != expected {
		t.Errorf("expcted: %f; actual: %f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	actual := Area(rectangle)
	expected := 100.0

	if actual != expected {
		t.Errorf("expcted: %f; actual: %f", expected, actual)
	}
}

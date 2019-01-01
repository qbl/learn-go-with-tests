package rectangle

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const tolerance = .01

var rectangle = Rectangle{10.0, 10.0}
var circle = Circle{7.0}

var opt = cmp.Comparer(func(x float64, y float64) bool {
	diff := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0
	return (diff / mean) < tolerance
})

func TestPerimeter(t *testing.T) {
	actual := Perimeter(rectangle)
	expected := 40.0

	if actual != expected {
		t.Errorf("expcted: %f; actual: %f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	t.Run("Area of a rectangle", func(t *testing.T) {
		actual := rectangle.Area()
		expected := 100.0

		if actual != expected {
			t.Errorf("expcted: %f; actual: %f", expected, actual)
		}
	})

	t.Run("Area of a circle", func(t *testing.T) {
		actual := circle.Area()
		expected := 153.93

		if !cmp.Equal(actual, expected, opt) {
			t.Errorf("expcted: %f; actual: %f", expected, actual)
		}
	})
}

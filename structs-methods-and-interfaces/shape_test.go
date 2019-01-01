package shape

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
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		if !cmp.Equal(actual, expected, opt) {
			t.Errorf("expcted: %f; actual: %f", expected, actual)
		}
	}

	t.Run("Area of a rectangle", func(t *testing.T) {
		checkArea(t, rectangle, 100.00)
	})

	t.Run("Area of a circle", func(t *testing.T) {
		checkArea(t, circle, 153.93)
	})
}

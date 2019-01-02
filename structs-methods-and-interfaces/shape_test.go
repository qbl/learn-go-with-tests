package shape

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const tolerance = .01

var rectangle = Rectangle{Width: 10.0, Height: 10.0}
var circle = Circle{Radius: 7.0}
var triangle = Triangle{Base: 12, Height: 6}

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
			t.Errorf("expcted: %.2f; actual: %.2f", expected, actual)
		}
	}

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{shape: rectangle, expected: 100.0},
		{shape: circle, expected: 153.93},
		{shape: triangle, expected: 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.expected)
	}
}

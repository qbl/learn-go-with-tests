package rectangle

import "math"

// Rectangle has width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle has radius
type Circle struct {
	Radius float64
}

// Perimeter accepts width and height, and returns perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area accepts widht and height, and returns area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area accepts radius, and returns area of a cirle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

package rectangle

type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter accepts width and height, and returns perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area accepts widht and height, and returns area of a rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

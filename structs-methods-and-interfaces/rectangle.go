package rectangle

// Perimeter accepts width and height, and returns perimeter of a rectangle
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// Area accepts widht and height, and returns area of a rectangle
func Area(width float64, height float64) float64 {
	return width * height
}

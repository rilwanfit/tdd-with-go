package structs

type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter which calculate the perimeter of a rectangle given a height and width.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area which returns the area of a rectangle.
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

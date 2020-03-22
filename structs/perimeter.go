package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Area which returns the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Redius float64
}

// Area which returns the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Redius * c.Redius
}

type Triangle struct {
	Base   float64
	Height float64
}

// Area which returns the area of a triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter which calculate the perimeter of a rectangle given a height and width.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

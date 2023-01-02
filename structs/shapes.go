package main

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of a given rectangle
func Perimeter(rectangle Rectangle) float64 {

	return (rectangle.Width + rectangle.Height) * 2
}

// Area calculates the area of a given rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

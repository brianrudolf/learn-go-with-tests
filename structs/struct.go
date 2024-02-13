package structs

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
	// return 2 * (width + height)
}

// func Area(rectangle Rectangle) float64 {
// 	// return width * height
// 	return rectangle.Width * rectangle.Height
// }

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Triangle) Area() float64 {
	return c.Base * c.Height * 0.5
}

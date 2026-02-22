package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}
type Circle struct {
	radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (t Triangle) Perimeter() float64 {
	return 2 * t.Base * t.Height
}
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

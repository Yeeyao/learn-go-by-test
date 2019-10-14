package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area(rectangle Rectangle) float64 {
	return r.Width * r.Height
}

func (c Circle) Area(circle Circle) float64 {
	return c.Radius * c.Radius * math.Pi
}

package main

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return .5 * t.Base * t. Height
}

type Shape interface {
	Area() float64
}

func Perimeter(shape Rectangle) float64 {
	total_width := shape.Width * 2
	total_height := shape.Height * 2

	return total_width + total_height
}

func Area(shape Rectangle) float64 {
	return shape.Width * shape.Height
}

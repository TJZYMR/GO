package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	length float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.length)
}

type Square struct {
	length float64
}

func (s Square) Area() float64 {
	return s.length * s.length
}

func (s Square) Perimeter() float64 {
	return 4 * s.length
}
func main() {
	var s Shape = Rectangle{width: 10, length: 20}
	fmt.Println(s.Area())
}

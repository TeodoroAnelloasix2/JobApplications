package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeData(s shape) string {

	text := fmt.Sprintf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
	return text
}

type shape interface {
	area() float64
	perimeter() float64
}

func main() {

	c := circle{radius: 5.5}
	r := rect{
		width:  4.3,
		height: 3,
	}
	println("Circulo: " + printShapeData(c))
	println("Rectangulo: " + printShapeData(r))
}

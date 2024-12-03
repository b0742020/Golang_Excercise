package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) float64 {
	return s.area()
}
func init() {
	fmt.Println("Initializing the program")
}

func main() {
	c1 := circle{
		radius: 5,
	}
	s1 := square{
		length: 10,
		width:  5,
	}
	fmt.Println("Area of square is:", info(s1))
	fmt.Println("Area of circle is:", info(c1))
}

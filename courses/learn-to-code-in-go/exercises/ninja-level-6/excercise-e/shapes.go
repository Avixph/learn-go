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

func (s square) area() float64 {
	area := math.Round(s.length * s.width)
	return area
}

func (c circle) area() float64 {
	area := math.Round(math.Pi * (c.radius * c.radius))
	return area
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh.area())
}

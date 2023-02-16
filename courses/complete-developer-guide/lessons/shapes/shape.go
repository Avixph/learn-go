package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (sq square) getArea() float64 {
	area := sq.sideLength * sq.sideLength
	return area
}

func (tr triangle) getArea() float64 {
	area := 0.5 * tr.base * tr.height
	return area
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

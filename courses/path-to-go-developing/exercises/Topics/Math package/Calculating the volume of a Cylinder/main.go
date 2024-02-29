package main

import (
	. "fmt"
	. "math"
)

func main() {
	var radius float64
	Scanln(&radius)

	var height float64
	Scanln(&height)

	// Using the functions from the math package create the formula
	// to calculate the volume of a cylinder below:
	volume := Pi * Pow(radius, 2) * height

	Printf("%.2f", volume)
}

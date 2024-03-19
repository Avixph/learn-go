package main

import (
	"fmt"
	"strconv"
)

const FloatPrecision = 64

func scanInput() float64 {
	var inp float64
	_, err := fmt.Scanln(&inp)
	if err != nil {
		return 0
	}
	return inp
}

func floatToString(flt float64) string {
	return strconv.FormatFloat(flt, 'f', -1, FloatPrecision)
}

func main() {
	//var mathConstant float64
	//fmt.Scanln(&mathConstant)
	mathConstant := scanInput()
	val := floatToString(mathConstant)
	fmt.Printf("%s", val)
}

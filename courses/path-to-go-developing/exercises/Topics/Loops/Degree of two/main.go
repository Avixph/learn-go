package main

import "fmt"

func getValue(x *int) {
	_, err := fmt.Scan(x)
	if err != nil {
		return
	}
}

func calculateDegree(x int) int {
	degree := 0
	value := 1

	for value < x {
		value *= 2
		degree++
	}

	return degree
}

func main() {
	var x int
	getValue(&x)
	fmt.Println(calculateDegree(x))
}

package main

import "fmt"

// Update the function's signature of getRectangleData() below:
func getRectangleData(length, width float64) (float64, float64) {
	area := length * width
	perimeter := 2 * (length + width)

	return area, perimeter
}

// DO NOT delete or modify the contents within the main function.
func main() {
	var length, width float64
	_, err := fmt.Scanln(&length, &width)
	if err != nil {
		return
	}

	area, perimeter := getRectangleData(length, width)
	fmt.Println("The area of the rectangle is:", area)
	fmt.Println("The perimeter of the rectangle is:", perimeter)
}

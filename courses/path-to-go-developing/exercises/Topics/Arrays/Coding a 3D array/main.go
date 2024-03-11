package main

import "fmt"

func main() {
	// Create the array and assign 88.6 to its [1][0][2] element below:
	var array [4][4][4]float32
	array[1][0][2] = 88.6

	// DO NOT delete, this line prints the array!
	fmt.Println(array)
}

package main

import "fmt"

func main() {
	// DO NOT delete the code block below!
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		return
	}

	// Use a nested for...range loop to print each element
	// of the matrix in a new line, multiplied by 'num'
	for _, row := range matrix {
		for _, column := range row {
			fmt.Println(column * num)
		}
	}
}

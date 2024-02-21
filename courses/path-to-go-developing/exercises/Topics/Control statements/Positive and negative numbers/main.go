package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d", &number)

	switch {
	case number < 0:
		fmt.Println("Negative!")
	case number > 0:
		fmt.Println("Positive!")
	default:
		fmt.Println("Zero!")
	}
}

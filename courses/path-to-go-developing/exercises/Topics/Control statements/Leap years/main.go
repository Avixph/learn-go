package main

import "fmt"

func main() {
	var year int
	fmt.Scanln(&year)

	switch {
	case year%4 == 0 && year%100 != 0 || year%400 == 0:
		fmt.Println("Leap year")
	default:
		fmt.Println("Regular year")
	}
}

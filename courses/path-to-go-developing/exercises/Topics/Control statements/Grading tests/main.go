package main

import "fmt"

// nolint:gomnd // DO NOT delete this comment!
func main() {
	var score int
	fmt.Scanf("%d", &score)

	switch {
	case score >= 71:
		fmt.Println("Passed!")
	default:
		fmt.Println("Failed!")
	}
}

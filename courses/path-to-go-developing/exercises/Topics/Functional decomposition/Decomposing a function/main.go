package main

import "fmt"

func f(x int) int {
	if x <= 0 {
		return f1(x)
	}
	return f2(x)
}

// nolint: gomnd // <- DO NOT delete this comment!
func f1(x int) int {
	return 2 * x
}

// nolint: gomnd // <- DO NOT delete this comment!
func f2(x int) int {
	return 3 * x
}

// DO NOT delete or modify the contents of the main function!
func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}

	result := f(n)
	fmt.Println(result)
}

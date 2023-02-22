package main

import "fmt"

func main() {
	fmt.Println("x =", x)
	fmt.Printf("x is of type %T\n", x)
	x = 42
	fmt.Println("x =", x)
}

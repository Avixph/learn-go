package main

import "fmt"

func main() {
	fmt.Println("x =", x)
	fmt.Printf("x is of type %T\n", x)
	x = 31
	fmt.Println("x =", x)
	y = int(x)
	fmt.Println("y =", y)
	fmt.Printf("y is of type %T\n", y)
}

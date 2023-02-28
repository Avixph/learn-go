package main

import "fmt"

func main() {
	point := true
	fmt.Println(point && true)
	fmt.Println(point && false)
	fmt.Println(point || true)
	fmt.Println(point || false)
	fmt.Println(!point)
}

package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("x = %v\ny = %v\nz = %v", x, y, z)
	fmt.Println(s)
}

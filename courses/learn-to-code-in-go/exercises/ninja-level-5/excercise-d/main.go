package main

import (
	"fmt"
)

func main() {
	nissaVersa := struct {
		doors  int
		color  string
		luxury bool
	}{
		doors:  4,
		color:  "red",
		luxury: false,
	}
	// print anon struct
	fmt.Println(nissaVersa.color)
}

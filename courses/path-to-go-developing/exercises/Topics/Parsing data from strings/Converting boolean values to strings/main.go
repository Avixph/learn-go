package main

import (
	"fmt"
	"strconv"
)

func scanInput() bool {
	var inp bool
	_, err := fmt.Scanln(&inp)
	if err != nil {
		return false
	}
	return inp
}

func main() {
	myBool := scanInput()
	val := strconv.FormatBool(myBool)
	fmt.Printf("%T %s", val, val)
}

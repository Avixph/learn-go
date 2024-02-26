package main

import (
	"errors"
	. "fmt"
	. "math"
)

func negativeNumError() error {
	// create a custom error message within the 'err' variable here
	err := errors.New("math: can't calculate square root of negative number")

	return err
}

// DO NOT delete or modify the code within the main() function!
func main() {
	var num float64
	_, err := Scanf("%f", &num)
	if err != nil {
		return
	}

	if num < 0 {
		err := negativeNumError()
		Println(err)
		return
	}

	Println(Sqrt(num))
}

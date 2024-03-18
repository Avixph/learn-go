package main

import (
    "fmt"
    "strings"
)

func main() {
	var builder strings.Builder

	var c byte  // what type allows us to read Unicode characters?
	fmt.Scanln(&c)

	for i := 0; i < 3; i++ {
		builder.WriteByte(c) // what function allows us to write Unicode characters?
	}

    // This line outputs the Unicode character do not delete it!
	fmt.Println(builder.String())
}
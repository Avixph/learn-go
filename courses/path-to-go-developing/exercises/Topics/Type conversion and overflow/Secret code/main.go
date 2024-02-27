package main

import "fmt"

func main() {
	var s string
	var shift int
	_, err := fmt.Scan(&s, &shift)
	if err != nil {
		return
	}

	for _, ch := range s {
		fmt.Print(string(ch + int32(shift))) // add the shift value to the character value
	}
}

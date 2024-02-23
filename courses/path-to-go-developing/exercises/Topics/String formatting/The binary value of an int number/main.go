package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		return
	}

	fmt.Printf("%b", num)
}

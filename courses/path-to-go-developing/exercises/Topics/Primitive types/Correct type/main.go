package main

import "fmt"

func main() {
	var number float32

	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}
	fmt.Println(number)
}

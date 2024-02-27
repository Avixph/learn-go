package main

import "fmt"

func main() {
	var v int

	// Put your code here
	_, err := fmt.Scan(&v)
	if err != nil {
		return
	}

	fmt.Println(v)
}

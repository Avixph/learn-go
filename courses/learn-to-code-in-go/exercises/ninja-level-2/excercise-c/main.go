package main

import "fmt"

func main() {
	// const total = 56.89
	// const item string = "fleece sweater"
	const (
		total        = 56.89
		item  string = "fleece sweater"
	)
	fmt.Println(total)
	fmt.Printf("%T\n", total)
	fmt.Println(item)
	fmt.Printf("%T\n", item)
}

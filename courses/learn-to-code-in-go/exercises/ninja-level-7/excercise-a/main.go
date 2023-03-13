package main

import "fmt"

func main() {
	name := "abf"
	address := &name
	fmt.Println(name)
	fmt.Println(&name)
	fmt.Println(*address)
}

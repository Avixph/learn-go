package main

import "fmt"

func main() {
	numA := foo(132)
	numB, str := bar(1891, "wow that's old!'")
	fmt.Println(numA, numB, str)
}

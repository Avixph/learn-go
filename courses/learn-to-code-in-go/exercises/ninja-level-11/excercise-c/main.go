package main

import "fmt"

func main() {
	e1 := customErr{
		info: "need more coffee",
	}

	foo(e1)
}

func foo(e error) {
	fmt.Println("foo ran -", e)
}

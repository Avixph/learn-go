package main

import "fmt"

func main() {
	for i := 10; i < 101; i++ {
		fmt.Printf("%v/4 is equal to %v with %v remaining.\n", i, (i / 4), i%4)
	}
}

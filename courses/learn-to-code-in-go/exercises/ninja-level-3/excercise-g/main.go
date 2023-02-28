package main

import "fmt"

func main() {
	for i := 1; i < 61; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else if !(i%3 != 0) {
			fmt.Printf("%v is odd\n", i)
		}
	}
}

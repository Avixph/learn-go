package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 110; i++ {
			if i%10 == 0 {
				c <- i
			}
			// c <- i
		}
		close(c)
	}()

	return c
}

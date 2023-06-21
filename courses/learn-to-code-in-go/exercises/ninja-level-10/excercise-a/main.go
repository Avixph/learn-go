package main

import "fmt"

func main() {
	c1 := make(chan int)

	// With annonymous function
	go func() {
		c1 <- 38
	}()

	// With buffer channel
	c2 := make(chan int, 1)
	c2 <- 76

	fmt.Println("your channel 1 --->", <-c1)
	fmt.Println("your channel 2 --->", <-c2)
}

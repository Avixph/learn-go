package main

import "fmt"

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

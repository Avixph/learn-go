package main

import "fmt"

var fooBar = func() {
	sum := 5
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}

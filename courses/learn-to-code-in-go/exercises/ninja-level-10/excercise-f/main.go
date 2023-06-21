package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		count := 100
		for i := 0; i < count; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("about to exit!")
}

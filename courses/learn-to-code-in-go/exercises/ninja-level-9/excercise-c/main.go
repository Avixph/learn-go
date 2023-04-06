package main

import (
	"fmt"
	"runtime"
)

func main() {
	counter := 0
	gr := 100
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Last value", counter)
}

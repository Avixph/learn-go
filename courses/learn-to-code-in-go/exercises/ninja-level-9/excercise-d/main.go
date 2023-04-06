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
			mx.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			mx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Last value", counter)
}

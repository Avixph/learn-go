package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	gr := 100
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			v := atomic.LoadInt64(&counter)
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Last value", counter)
}

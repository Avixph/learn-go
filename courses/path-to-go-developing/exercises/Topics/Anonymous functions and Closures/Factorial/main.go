package main

import "fmt"

func main() {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}

	factorial := func() int {
		f := 1
		for i := 1; i <= number; i++ {
			f *= i
		}
		return f
	}()

	fmt.Println(factorial)
}

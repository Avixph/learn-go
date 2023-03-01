package main

import "fmt"

func main() {
	nums := []int{6, 12, 24, 36, 48, 60, 72, 84, 96, 120}
	for _, num := range nums {
		fmt.Println(num)
		fmt.Printf("num := %T\n", num)
	}
	fmt.Printf("nums := %T\n", nums)
}

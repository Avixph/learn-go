package main

import "fmt"

func main() {
	birthDay := 1991
	// for i := birthDay; i < 2023; i++ {
	// 	fmt.Println(i)
	// }
	for birthDay <= 2023 {
		fmt.Println(birthDay)
		birthDay++
	}
}

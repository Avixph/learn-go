package main

import "fmt"

func factorial(num int) int {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	return res
}

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	fmt.Println(factorial(num))
}

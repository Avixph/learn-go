package main

import "fmt"

func checkNum() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		evenOrOdd(num)
	}
}

func evenOrOdd(num int) {
	isEven := fmt.Sprintf("%v is an even number.", num)
	isOdd := fmt.Sprintf("%v is an odd number.", num)
	if num%2 == 0 {
		fmt.Println(isEven)
	} else {
		fmt.Println(isOdd)
	}
}

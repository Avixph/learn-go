package main

import "fmt"

func reverseNum(num int) int {
	rev := 0
	for num > 0 {
		rev = rev*10 + num%10
		num = num / 10
	}
	return rev
}

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	fmt.Println(reverseNum(num))
}

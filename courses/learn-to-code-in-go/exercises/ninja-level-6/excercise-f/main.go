package main

import "fmt"

func main() {
	nums := []int{8, 16, 24, 32}
	sum := 0

	func(nums []int) {
		for _, num := range nums {
			// fmt.Println(num)
			sum += num
		}
	}(nums)
	fmt.Println(sum)
}

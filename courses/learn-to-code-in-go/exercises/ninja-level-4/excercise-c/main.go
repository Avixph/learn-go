package main

import "fmt"

func main() {
	nums := []int{6, 12, 24, 36, 48, 60, 72, 84, 96, 120}
	fmt.Println(nums[:5])
	fmt.Println(nums[5:])
	fmt.Println(nums[2:7])
	fmt.Println(nums[1:6])
}

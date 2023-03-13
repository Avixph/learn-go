package main

import "fmt"

func main() {
	nums := []int{91, 92, 93, 94, 95, 96, 97, 99}
	sum := fooBar(nums...)
	evenSum := even(fooBar, nums...)
	oddSum := odd(fooBar, nums...)
	fmt.Println("Sum of all nums is", sum)
	fmt.Println("Sum of all even nums is", evenSum)
	fmt.Println("Sum of all odd nums is", oddSum)
}

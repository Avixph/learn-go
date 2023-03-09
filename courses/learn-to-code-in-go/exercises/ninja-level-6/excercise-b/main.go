package main

import "fmt"

func main() {
	nums1 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}
	nums2 := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29}
	fmt.Println("sum of nums1 is", foo(nums1...))
	fmt.Println("sum of nums2 is", bar(nums2))
}

package main

func foo(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func bar(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

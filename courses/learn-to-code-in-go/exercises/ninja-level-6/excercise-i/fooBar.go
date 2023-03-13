package main

func fooBar(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func even(fun func(nums ...int) int, nums ...int) int {
	var evens []int
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return fun(evens...)
}

func odd(fun func(nums ...int) int, nums ...int) int {
	var odds []int
	for _, num := range nums {
		if num%2 != 0 {
			odds = append(odds, num)
		}
	}
	return fun(odds...)
}

package main

func fooBar() func() []int {
	var nums []int
	for i := 0; i < 25; i++ {
		if i%2 == 0 {
			nums = append(nums, i)
		}
	}
	return func() []int {
		return nums
	}
}

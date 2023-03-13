package main

func fooBar() func() int {
	x := 0
	return func() int {
		for i := 0; i < 25; i++ {
			if i%2 == 0 {
				x += i
			}
		}
		return x
	}
}

package main

import . "fmt"

func getNumber() int {
	var num int
	_, err := Scan(&num)
	if err != nil {
		return 0
	}
	return num
}

func intToBool(num int) bool {
	res := num != 0
	return res
}

func main() {
	result := intToBool(getNumber())
	Println(result)
}

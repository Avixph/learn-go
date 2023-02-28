package main

import "fmt"

func main() {
	birthDay := 1991
	for {
		if birthDay <= 2023 {
			fmt.Println(birthDay)
			birthDay++
		} else {
			break
		}
	}
}

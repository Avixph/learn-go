package main

import "fmt"

func main() {
	var num1, num2 int8
	_, err := fmt.Scan(&num1, &num2)
	if err != nil {
		return
	}

	if int(num1+num2) != int(num1)+int(num2) {
		fmt.Println("+")
	}
	if int(num1-num2) != int(num1)-int(num2) {
		fmt.Println("-")
	}
	if int(num1*num2) != int(num1)*int(num2) {
		fmt.Println("*")
	}
}

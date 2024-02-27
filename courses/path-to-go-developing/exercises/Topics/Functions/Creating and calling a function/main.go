package main

import "fmt"

func userInfo(user string) string {
	return fmt.Sprintf("%s is learning how to call functions!", user)
}

func main() {
	// Do not change this two lines of code
	var userName string
	_, err := fmt.Scanf("%s", &userName)
	if err != nil {
		return
	}

	fmt.Println(userInfo(userName))
}

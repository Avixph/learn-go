package main

import (
	"fmt"
	"net/http"
)

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	var message string
	if err != nil {
		message = fmt.Sprintf("%v may be down at the momment!", err)
		fmt.Println(message)
		c <- url
		return
	}
	message = fmt.Sprintf("%v is up and running!", url)
	fmt.Println(message)
	c <- url

}

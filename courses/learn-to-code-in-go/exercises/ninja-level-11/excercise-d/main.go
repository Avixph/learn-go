package main

import "log"

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

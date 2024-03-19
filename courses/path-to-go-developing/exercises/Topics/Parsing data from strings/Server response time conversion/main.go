package main

import (
	"fmt"
	"log"
	"strconv"
)

func scanInput() string {
	var inp string
	_, err := fmt.Scanln(&inp)
	if err != nil {
		return ""
	}
	return inp
}

func main() {
	//var logEntry string
	//fmt.Scanln(&logEntry)
	logEntry := scanInput()
	logEntry = logEntry[:len(logEntry)-2]

	responseTime, err := strconv.ParseFloat(logEntry, 64)
	if err != nil {
		log.Fatal(err)
	}

	responseTime /= 1000

	fmt.Printf("Response time: %.2f seconds", responseTime)
}

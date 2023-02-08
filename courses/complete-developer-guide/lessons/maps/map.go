package main

import "fmt"

func printMap(cMap map[string]string) {
	for color, hex := range cMap {
		message := fmt.Sprintf("The hex-code for %v is %v.", color, hex)
		fmt.Println(message)
	}

}

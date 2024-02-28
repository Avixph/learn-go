package main

import "fmt"

func main() {
	// DO NOT delete or modify the code block below, it reads a random input word
	var word string
	_, err := fmt.Scan(&word)
	if err != nil {
		return
	}

	// Use slice expressions with the len() function
	// to remove all the letters from 'word' except for the last letter.
	word = word[len(word)-1:]

	fmt.Println(word)
}

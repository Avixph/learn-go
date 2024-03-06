package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("keanu_quotes.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	quotes := []string{"You're breathtaking!",
		"Wake up Samurai! We have a city to burn.",
		"Lose? I don't lose; I win! I'm a lawyer, that's my job, that's what I do!",
	}

	for _, line := range quotes {
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			log.Fatal(err)
		}
	}
}

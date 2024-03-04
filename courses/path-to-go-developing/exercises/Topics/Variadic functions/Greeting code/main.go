package main

import "fmt"

func main() {
	var prefix, name1, name2 string
	_, err := fmt.Scan(&prefix, &name1, &name2)
	if err != nil {
		return
	}

	for _, line := range Greeting(prefix, name1, name2) {
		fmt.Println(line)
	}
}

func Greeting(prefix string, names ...string) []string {
	var greet []string
	for _, name := range names {
		greet = append(greet, fmt.Sprintf("%s %s", prefix, name))
	}
	return greet
}

package main

import "fmt"

func main() {
	jbCharacters := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Woman"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "unsets"},
	}

	jbCharacters["fleming_ian"] = []string{"steak", "cigars", "espionage"}

	for i, characters := range jbCharacters {
		fmt.Println("character:", i)
		for j, character := range characters {
			fmt.Println(j+1, character)
		}
	}
}
